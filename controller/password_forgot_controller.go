package controller

import (
	"html/template"
	"net/http"
)

func PasswordForgotController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/pass-forgot.html")
		_ = t.Execute(w, nil)
	}
}
