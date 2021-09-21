package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func main() {
	_ = template.Must(template.ParseFiles("templates/login.gohtml"))
	http.HandleFunc("/login", handler)
	http.ListenAndServe(":8085", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/login.gohtml"))
	if r.Method == "GET" {
		t.ExecuteTemplate(w, "login.gohtml", nil)
	}

	if r.Method == "POST" {

		r.ParseForm()
		username := strings.Join(r.Form["username"], "")
		fmt.Print(username)
		if username == "" {
			t.ExecuteTemplate(w, "login.gohtml", "Username  doesn't empty")
			return
		}

	}
}
