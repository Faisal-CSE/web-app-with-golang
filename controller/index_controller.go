package controller

import (
	"html/template"
	"net/http"
)

func IndexController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/index.html")
		err := t.Execute(w, nil)
		if err != nil {
			return
		}
	} else {
		_, err := w.Write([]byte("<h2>Invalid Method!</h2>"))
		if err != nil {
			return
		}
	}
}
