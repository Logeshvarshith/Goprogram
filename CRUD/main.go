package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	Id    int64
	Fname string
	Lname string
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handler)
	http.HandleFunc("/create", create)
	http.HandleFunc("/view", shows)
	http.HandleFunc("/process", processor)

	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/main.gohtml"))

	t.ExecuteTemplate(w, "main.gohtml", nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/create.gohtml"))
	t.ExecuteTemplate(w, "create.gohtml", nil)
}
func shows(w http.ResponseWriter, r *http.Request) {
	rollno, _ := strconv.ParseInt(r.FormValue("rollno"), 0, 0)
	fname := r.FormValue("firster")
	lname := r.FormValue("laster")
	t := template.Must(template.ParseFiles("templates/view.gohtml"))
	d := struct {
		Rollno int64
		First  string
		Last   string
	}{
		Rollno: rollno,
		First:  fname,
		Last:   lname,
	}
	var err error
	db, err := sql.Open("mysql", "root:Root@12345@tcp(127.0.0.1:3306)/sample")
	if err != nil {
		log.Fatal(err)
	}
	pingerr := db.Ping()
	if pingerr != nil {
		log.Fatal(pingerr)
	}

	_, err = db.Exec("INSERT INTO sample.check (id, fname, lname) VALUES (?, ?, ?)", d.Rollno, d.First, d.Last)
	if err != nil {
		log.Fatal(err)
	}

	t.ExecuteTemplate(w, "view.gohtml", d)
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		//log.Fatal(r.Method)

	}

	t := template.Must(template.ParseFiles("templates/update.gohtml"))
	rollno, _ := strconv.ParseInt(r.FormValue("rollno"), 0, 0)
	var err error
	db, err := sql.Open("mysql", "root:Root@12345@tcp(127.0.0.1:3306)/sample")
	if err != nil {
		log.Fatal(err)
	}
	pingerr := db.Ping()
	if pingerr != nil {
		log.Fatal(pingerr)
	}
	var a student
	row := db.QueryRow("SELECT * FROM sample.check WHERE id = ?", rollno)
	row.Scan(&a.Id, &a.Fname, &a.Lname)

	d := struct {
		Rollno int64
		First  string
		Last   string
	}{
		Rollno: a.Id,
		First:  a.Fname,
		Last:   a.Lname,
	}

	t.ExecuteTemplate(w, "update.gohtml", d)
}
