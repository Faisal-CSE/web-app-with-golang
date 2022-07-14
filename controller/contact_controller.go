package controller

import (
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func ContactController(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/contact_form.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	details := ContactDetails{
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}

	// do something with details
	_ = details

	tmpl.Execute(w, struct{ Success bool }{true})
}
