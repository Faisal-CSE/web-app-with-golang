package main

import (
	"github.com/gorilla/mux"
	v1API "github.com/web-app-with-golang/api/v1"
	"github.com/web-app-with-golang/controller"
	env "github.com/web-app-with-golang/project_env"
	"log"
	"net/http"
)

func main() {
	var strAPIVersion1 string

	port := env.GoDotEnvVariable("PORT")
	if port == "" {
		port = "8080"
	}

	route := mux.NewRouter()

	////route := http.NewServeMux()
	//// Create a file server which serves files out of the "./assets" directory.
	//// Note that the path given to the http.Dir function is relative to the project
	//// directory root.
	//fs := http.FileServer(http.Dir("./assets/"))
	//
	//// Use the mux.Handle() function to register the file server as the handler for
	//// all URL paths that start with "/static/". For matching paths, we strip the
	//// "/static" prefix before the request reaches the file server.
	//route.Handle("/static/", http.StripPrefix("/static", fs))

	route.HandleFunc("/", controller.IndexController).Methods(http.MethodGet)
	route.HandleFunc("/login", controller.LoginController)
	route.HandleFunc("/contact", controller.ContactController)
	route.HandleFunc("/test", controller.DummyController)

	// API URL VERSION 1
	strAPIVersion1 = "/api/v1/"
	route.HandleFunc(strAPIVersion1+"sample-health-check", v1API.SampleHealthCheckAPI).Methods(http.MethodGet)

	log.Print("Starting server on :" + port)
	_ = http.ListenAndServe(":"+port, route)
}
