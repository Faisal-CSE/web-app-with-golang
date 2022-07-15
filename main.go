package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	v1API "github.com/web-app-with-golang/api/v1"
	"github.com/web-app-with-golang/controller"
	"log"
	"net/http"
	"os"
)

func main() {
	var strAPIVersion1 string
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	route := mux.NewRouter()

	route.HandleFunc("/", controller.IndexController).Methods("GET")
	route.HandleFunc("/login", controller.LoginController)
	route.HandleFunc("/contact", controller.ContactController)
	route.HandleFunc("/test", controller.DummyController)

	// API URL VERSION 1
	strAPIVersion1 = "/api/v1/"
	route.HandleFunc(strAPIVersion1+"sample-health-check", v1API.SampleHealthCheckAPI).Methods("GET")

	errListener := http.ListenAndServe(":"+port, route)
	if errListener != nil {
		return
	}
}
