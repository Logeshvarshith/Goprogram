package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

type TempData struct {
	Username string
	Email    string
	AuthInfo string
}

func main() {
	http.HandleFunc("/forgot", forgotpasswordHandler)
	http.HandleFunc("/forgotauth", forgotpasswordauthhandler)
	http.HandleFunc("/forgotpwchange", forgotpwChangeHandler) // renders change pw form and places authInfo in form action
	http.HandleFunc("/forgotpwemailver", forgotPWverHandler)

	http.ListenAndServe(":8086", nil)

}

func forgotpasswordHandler(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/forgotpassword.html"))
	t.ExecuteTemplate(w, "forgotpassword.html", nil)

}

/*func forgotpasswordauthhandler(w http.ResponseWriter, r *http.Request) {
	var err error
	db, err = sql.Open("mysql", "root:Root@12345@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	t := template.Must(template.ParseFiles("templates/forgotpassword.html"))
	r.ParseForm()
	email := r.FormValue("email")

	stmt := "SELECT email,Username FROM testdb.bcrypt WHERE email = ?"
	row := db.QueryRow(stmt, email)
	var Username string
	err = row.Scan(&email, &Username)
	if err != nil {
		fmt.Println("username not exists, err:", err)
		t.ExecuteTemplate(w, "forgotpassword.html", "Username not present")
		return
	}

	now := time.Now()

	timeout := now.Add(time.Minute * 45)

	rand.Seed(time.Now().UnixNano())

	var alphaNumRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	emailVerRandRune := make([]rune, 64)
	for i := 0; i < 64; i++ {
		// Intn() returns, as an int, a non-negative pseudo-random number in [0,n).
		emailVerRandRune[i] = alphaNumRunes[rand.Intn(len(alphaNumRunes)-1)]
	}

	emailVerPassword := string(emailVerRandRune)

	var emailVerPWhash []byte
	emailVerPWhash, err = bcrypt.GenerateFromPassword([]byte(emailVerPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt err:", err)
		t.ExecuteTemplate(w, "forgotpassword.html", "Email was not sent")
		return
	}
	var updateEmailVerStmt *sql.Stmt
	updateEmailVerStmt, err = db.Prepare("UPDATE email_ver_hash SET ver_hash = ?, timeout = ? WHERE email = ?;")
	if err != nil {
		fmt.Println("error preparing statement:", err)

		t.ExecuteTemplate(w, "forgotpassword.html", "db update error occured")
		return
	}

	emailVerPWhashStr := string(emailVerPWhash)
	fmt.Print(emailVerPWhashStr)

	result, _ := updateEmailVerStmt.Exec(emailVerPWhashStr, timeout, email)
	fmt.Print(result)
	os.Setenv("FromEmailAddr", "logesh1104@gmail.com")
	os.Setenv("SMTPpwd", "15ec046logesh")
	os.Setenv("ToEmailAddr", email)
	// sender data
	from := os.Getenv("FromEmailAddr") //ex: "John.Doe@gmail.com"
	password := os.Getenv("SMTPpwd")
	// receiver address
	toEmail := os.Getenv("ToEmailAddr") // ex: "Jane.Smith@yahoo.com"
	to := []string{toEmail}
	// smtp - Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	// message
	subject := "Subject:Password Request Mail\n"
	body := "<body><a rel=\"nofollow noopener noreferrer\" target=\"_blank\" href=\"https://www.mysite.com/forgotpwchange?u=" + Username + "&evpw=" + emailVerPassword + "\">Change Password</a></body>"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + body)
	// athentication data
	// func PlainAuth(identity, username, password, host string) Auth
	auth := smtp.PlainAuth("", from, password, host)
	// send mail
	// func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
	err = smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Successfully sent mail to all user in toList")
	t = template.Must(template.ParseFiles("templates/recovery.html"))

	t.ExecuteTemplate(w, "recovery.html", "Your password is send successfully")

}*/
func forgotpasswordauthhandler(w http.ResponseWriter, r *http.Request) {
	var err error
	db, err = sql.Open("mysql", "root:Root@12345@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	t := template.Must(template.ParseFiles("templates/forgotpassword.html"))
	r.ParseForm()
	email := r.FormValue("email")
	username := r.FormValue("username")

	stmt := "SELECT email,Username FROM testdb.bcrypt WHERE email = ? AND Username=?"
	row := db.QueryRow(stmt, email, username)
	err = row.Scan(&email, &username)
	if err != nil {
		fmt.Println("username not exists, err:", err)
		t.ExecuteTemplate(w, "forgotpassword.html", "Username and email not present")
		return
	}

	os.Setenv("FromEmailAddr", "logesh1104@gmail.com")
	os.Setenv("SMTPpwd", "15ec046logesh")
	os.Setenv("ToEmailAddr", email)
	// sender data
	from := os.Getenv("FromEmailAddr") //ex: "John.Doe@gmail.com"
	password := os.Getenv("SMTPpwd")
	// receiver address
	toEmail := os.Getenv("ToEmailAddr") // ex: "Jane.Smith@yahoo.com"
	to := []string{toEmail}
	// smtp - Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	a := "%329898920q90q9"
	// message
	subject := "Subject:Password Request Mail\n"
	body := "<body><h>Click the below link to change the password </h><a rel=\"nofollow noopener noreferrer\" target=\"_blank\" href=\"http://localhost:8086/forgotpwchange?u=" + username + "&evpw=" + a + "\">Change Password</a></body>"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + body)
	// athentication data
	// func PlainAuth(identity, username, password, host string) Auth
	auth := smtp.PlainAuth("", from, password, host)
	// send mail
	// func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
	err = smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Successfully sent mail to all user in toList")
	t = template.Must(template.ParseFiles("templates/recovery.html"))

	t.ExecuteTemplate(w, "recovery.html", "Your password is send successfully")

}
func forgotpwChangeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****forgotpwChangeHandler running*****")
	t := template.Must(template.ParseFiles("templates/forgotpwchange.html"))
	username := r.FormValue("u")
	fmt.Println("username:", username)

	var td TempData
	td.AuthInfo = "?u=" + username
	t.ExecuteTemplate(w, "forgotpwchange.html", td)
}
func forgotPWverHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseGlob("templates/*.html")
	fmt.Println("*****forgotPWverHandler running*****")
	username := r.FormValue("u")
	userPassword := r.FormValue("password")
	confirmPassword := r.FormValue("confirmpassword")
	fmt.Println("username:", username)
	fmt.Println("userPassword:", userPassword)
	fmt.Println("confirmPassword:", confirmPassword)
	if userPassword == "" && confirmPassword == "" {
		fmt.Println("passwords not empty")

		t.ExecuteTemplate(w, "recovery.html", "Passwords not empty")
		return
	}

	if userPassword != confirmPassword {
		fmt.Println("passwords do no match")

		t.ExecuteTemplate(w, "recovery.html", "Passwords do not match")
		return
	}
	db, _ := sql.Open("mysql", "root:Root@12345@tcp(127.0.0.1:3306)/testdb")

	// check if db ver_hash is the same as the hash of emailVerPassword from email

	// check userPassword criteria
	// generate hash for new userPassword
	var hash []byte
	// generate emailVerPassword hash for db
	hash, _ = bcrypt.GenerateFromPassword([]byte(confirmPassword), bcrypt.DefaultCost)
	// update db with new userPasswordHash
	stmt := "SELECT Hash FROM testdb.bcrypt WHERE  Username=?"
	row := db.QueryRow(stmt, username)
	var hashed []byte
	err := row.Scan(&hashed)
	if err != nil {
		fmt.Print(hashed)
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(confirmPassword))
	if err == nil {
		t.ExecuteTemplate(w, "recovery.html", "Password does not match with previous password.Please fill the new password")
		return
	}

	stmt = "UPDATE bcrypt SET Hash = ? WHERE Username = ?"
	updateHashStmt, err := db.Prepare(stmt)
	if err != nil {

		t.ExecuteTemplate(w, "forgotpassword.html", "db update failed")
		return
	}
	defer updateHashStmt.Close()
	var result sql.Result
	result, err = updateHashStmt.Exec(hash, username)
	rowsAff, _ := result.RowsAffected()
	fmt.Println("rowsAff:", rowsAff)
	// check for successfull insert
	if err != nil || rowsAff != 1 {

		t.ExecuteTemplate(w, "forgotpassword.html", "db insert failed")
		return
	}
	t.ExecuteTemplate(w, "recovery.html", "password reset succesfully")
}
