package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

func LoginController(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/login_form.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}
