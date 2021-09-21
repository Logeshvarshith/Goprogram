package main

import (
	"database/sql"
	"fmt"
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
	http.HandleFunc("/process", processor)

	http.ListenAndServe(":8082", nil)
	details, err := show(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", details)

}

func handler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/main.gohtml"))

	t.ExecuteTemplate(w, "main.gohtml", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		//log.Fatal(r.Method)

	}

	t := template.Must(template.ParseFiles("templates/processor.gohtml"))
	rollno, _ := strconv.ParseInt(r.FormValue("rollno"), 0, 0)
	fname := r.FormValue("firster")
	lname := r.FormValue("laster")

	print(rollno, fname, lname)

	d := struct {
		Rollno int64
		First  string
		Last   string
	}{
		Rollno: rollno,
		First:  fname,
		Last:   lname,
	}

	t.ExecuteTemplate(w, "processor.gohtml", d)
}
func print(id int64, a string, b string) {
	var err error
	db, err := sql.Open("mysql", "root:Root@12345@tcp(127.0.0.1:3306)/sample")
	if err != nil {
		log.Fatal(err)
	}
	pingerr := db.Ping()
	if pingerr != nil {
		log.Fatal(pingerr)
	}
	_, err = db.Exec("INSERT INTO sample.check (id,fname, lname) VALUES (?,?, ?)", id, a, b)

	if err != nil {
		log.Fatal(err)
	}
}
func show(id int64) (student, error) {
	db, err := sql.Open("mysql", "root:Root@123456@tcp(127.0.0.1:3306)/sample")
	if err != nil {
		log.Fatal(err)
	}
	pingerr := db.Ping()
	if pingerr != nil {
		log.Fatal(pingerr)
	}
	var a student
	row := db.QueryRow("SELECT * FROM sample.check WHERE id = ?", id)
	row.Scan(&a.Id, &a.Fname, &a.Lname)
	return a, nil

}
