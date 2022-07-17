package controller

import (
	"net/http"
	"text/template"
)

func RegistrationController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/register.html")
		_ = t.Execute(w, nil)
	}
}
