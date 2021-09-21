package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"unicode"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
)

var tpl *template.Template
var db *sql.DB

func main() {
	tpl, _ = template.ParseGlob("templates/*.html")
	var err error
	db, err = sql.Open("mysql", "root:Root@12345@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/loginauth", loginAuthHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/registerauth", registerAuthHandler)
	http.HandleFunc("/password", passwordHandler)
	http.HandleFunc("/passwordauth", passwordauthhandler)

	http.ListenAndServe(":8080", nil)
	fmt.Print("Listen")
}

// loginHandler serves form for users to login with
func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****loginHandler running*****")
	tpl.ExecuteTemplate(w, "login.html", nil)
}

// loginAuthHandler authenticates user login
func loginAuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****loginAuthHandler running*****")
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println("username:", username, "password:", password)
	// retrieve password from db to compare (hash) with user supplied password's hash
	var hash string
	stmt := "SELECT Hash FROM bcrypt WHERE Username = ?"
	row := db.QueryRow(stmt, username)
	fmt.Print(row)
	fmt.Print(&hash)
	err := row.Scan(&hash)
	fmt.Println("hash from db:", hash)
	if err != nil {
		fmt.Println("error selecting Hash in db by Username")
		tpl.ExecuteTemplate(w, "login.html", "check username and password")
		return
	}
	// func CompareHashAndPassword(hashedPassword, password []byte) error
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	// returns nill on succcess
	if err == nil {
		tpl.ExecuteTemplate(w, "display.html", "You have successfully logged in :)")
		return
	}
	fmt.Println("incorrect password")
	tpl.ExecuteTemplate(w, "login.html", "check username and password")
}

// registerHandler serves form for registring new users
func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****registerHandler running*****")
	tpl.ExecuteTemplate(w, "register.html", nil)
}

// registerAuthHandler creates new user in database
func registerAuthHandler(w http.ResponseWriter, r *http.Request) {
	/*
		1. check username criteria
		2. check password criteria
		3. check if username is already exists in database
		4. create bcrypt hash from password
		5. insert username and password hash in database
		(email validation will be in another video)
	*/
	fmt.Println("*****registerAuthHandler running*****")
	r.ParseForm()
	username := r.FormValue("username")
	email := r.FormValue("email")
	// check username for only alphaNumeric characters
	var nameAlphaNumeric = true
	for _, char := range username {
		// func IsLetter(r rune) bool, func IsNumber(r rune) bool
		// if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			nameAlphaNumeric = false
		}
	}
	// check username pswdLength
	var nameLength bool
	if 4 <= len(username) && len(username) <= 50 {
		nameLength = true
	}
	// check password criteria
	password := r.FormValue("password")
	fmt.Println("password:", password, "\npswdLength:", len(password))
	// variables that must pass for password creation criteria
	var pswdLowercase, pswdUppercase, pswdNumber, pswdSpecial, pswdLength, pswdNoSpaces bool
	pswdNoSpaces = true
	for _, char := range password {
		switch {
		// func IsLower(r rune) bool
		case unicode.IsLower(char):
			pswdLowercase = true
		// func IsUpper(r rune) bool
		case unicode.IsUpper(char):
			pswdUppercase = true
		// func IsNumber(r rune) bool
		case unicode.IsNumber(char):
			pswdNumber = true
		// func IsPunct(r rune) bool, func IsSymbol(r rune) bool
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			pswdSpecial = true
		// func IsSpace(r rune) bool, type rune = int32
		case unicode.IsSpace(int32(char)):
			pswdNoSpaces = false
		}
	}
	if 5 < len(password) && len(password) < 15 {
		pswdLength = true
	}
	fmt.Println("pswdLowercase:", pswdLowercase, "\npswdUppercase:", pswdUppercase, "\npswdNumber:", pswdNumber, "\npswdSpecial:", pswdSpecial, "\npswdLength:", pswdLength, "\npswdNoSpaces:", pswdNoSpaces, "\nnameAlphaNumeric:", nameAlphaNumeric, "\nnameLength:", nameLength)
	if !pswdLowercase || !pswdUppercase || !pswdNumber || !pswdSpecial || !pswdLength || !pswdNoSpaces || !nameAlphaNumeric || !nameLength {
		tpl.ExecuteTemplate(w, "register.html", "please check username and password criteria")
		return
	}
	// check if username already exists for availability
	stmt := "SELECT UserID FROM bcrypt WHERE username = ?"
	row := db.QueryRow(stmt, username)
	var uID string
	err := row.Scan(&uID)
	if err != sql.ErrNoRows {
		fmt.Println("username already exists, err:", err)
		tpl.ExecuteTemplate(w, "register.html", "username already taken")
		return
	}
	// create hash from password
	var hash []byte
	// func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt err:", err)
		tpl.ExecuteTemplate(w, "register.html", "there was a problem registering account")
		return
	}
	fmt.Println("hash:", hash)
	fmt.Println("string(hash):", string(hash))
	// func (db *DB) Prepare(query string) (*Stmt, error)
	var insertStmt *sql.Stmt
	insertStmt, err = db.Prepare("INSERT INTO bcrypt (Username, Hash,email) VALUES (?, ?,?);")
	if err != nil {
		fmt.Println("error preparing statement:", err)
		tpl.ExecuteTemplate(w, "register.html", "there was a problem registering account")
		return
	}
	defer insertStmt.Close()
	var result sql.Result
	//  func (s *Stmt) Exec(args ...interface{}) (Result, error)
	result, err = insertStmt.Exec(username, hash, email)
	rowsAff, _ := result.RowsAffected()
	lastIns, _ := result.LastInsertId()
	fmt.Println("rowsAff:", rowsAff)
	fmt.Println("lastIns:", lastIns)
	fmt.Println("err:", err)
	if err != nil {
		fmt.Println("error inserting new user")
		tpl.ExecuteTemplate(w, "register.html", "there was a problem registering account")
		return
	}
	tpl.ExecuteTemplate(w, "success.html", "congrats, your account has been successfully created")
}
func passwordHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/forgotpwd.html"))
	fmt.Println("*****PasswordHandler running*****")
	t.ExecuteTemplate(w, "forgotpwd.html", nil)
}
func passwordauthhandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/forgotpwd.html"))
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")
	fmt.Print(email)
	stmt := "SELECT UserID FROM bcrypt WHERE username = ?"
	row := db.QueryRow(stmt, username)
	var uID string
	err := row.Scan(&uID)
	if err == sql.ErrNoRows {
		fmt.Println("username not exists, err:", err)
		t.ExecuteTemplate(w, "forgotpwd.html", "Username not present")
		return
	}
	stmt = "SELECT email FROM bcrypt WHERE username=?"
	row = db.QueryRow(stmt, username)
	var gmail string
	err = row.Scan(&gmail)
	if err != nil {
		fmt.Println("bcrypt err:", err)
		tpl.ExecuteTemplate(w, "forgotpwd.html", "Please enter the correct email id")
		return
	}
	var hash []byte
	// func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt err:", err)
		tpl.ExecuteTemplate(w, "forgotpwd.html", "there was a problem in reset password")
		return
	}
	// check password criteria
	password = r.FormValue("password")
	confirmpassword := r.FormValue("cpassword")
	if password != confirmpassword {
		tpl.ExecuteTemplate(w, "forgotpwd.html", "please check the password")
		return
	}
	stmt = "SELECT Hash FROM bcrypt WHERE username=?"
	row = db.QueryRow(stmt, username)
	var hashed []byte
	err = row.Scan(&hashed)
	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err == nil {
		tpl.ExecuteTemplate(w, "forgotpwd.html", "Password does not match with previous password")
		return
	}

	fmt.Println("password:", password, "\npswdLength:", len(password))
	// variables that must pass for password creation criteria
	var pswdLowercase, pswdUppercase, pswdNumber, pswdSpecial, pswdLength, pswdNoSpaces bool
	pswdNoSpaces = true
	for _, char := range password {
		switch {
		// func IsLower(r rune) bool
		case unicode.IsLower(char):
			pswdLowercase = true
		// func IsUpper(r rune) bool
		case unicode.IsUpper(char):
			pswdUppercase = true
		// func IsNumber(r rune) bool
		case unicode.IsNumber(char):
			pswdNumber = true
		// func IsPunct(r rune) bool, func IsSymbol(r rune) bool
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			pswdSpecial = true
		// func IsSpace(r rune) bool, type rune = int32
		case unicode.IsSpace(int32(char)):
			pswdNoSpaces = false
		}
	}
	if 5 < len(password) && len(password) < 15 {
		pswdLength = true
	}
	fmt.Println("pswdLowercase:", pswdLowercase, "\npswdUppercase:", pswdUppercase, "\npswdNumber:", pswdNumber, "\npswdSpecial:", pswdSpecial, "\npswdLength:", pswdLength, "\npswdNoSpaces:", pswdNoSpaces)
	fmt.Print(!pswdLowercase)
	fmt.Print(!pswdNoSpaces)
	if !pswdLowercase || !pswdUppercase || !pswdNumber || !pswdSpecial || !pswdLength || !pswdNoSpaces {
		tpl.ExecuteTemplate(w, "forgotpwd.html", "please check username and password criteria")
		return
	}
	user := r.FormValue("username")
	//var tx *sql.Tx
	db, _ = sql.Open("mysql", "root:Root@12345@tcp(127.0.0.1:3306)/testdb")
	db.Ping()
	new, err := db.Exec("UPDATE testdb.bcrypt SET hash = ? WHERE Username = ?", hash, user)
	if err != nil {
		fmt.Print(err, new)
		tpl.ExecuteTemplate(w, "forgotpwd.html", "there was a problem in reset password")
		fmt.Println("bcrypt err:", err)
		return
	}
	tpl.ExecuteTemplate(w, "forgotsuccess.html", "congrats, your password reset successfully")
	/*	if err == nil {
		t := template.Must(template.ParseFiles("templates/forgotsuccess.html"))
		t.ExecuteTemplate(w, "forgotsuccess.html", "Password Reset Successfully :)")
		return
	}*/

}

//forgot password email link
