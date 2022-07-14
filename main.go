package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/web-app-with-golang/controller"
	"log"
	"net/http"
	"os"
)

func main() {
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
	route.HandleFunc("/login", controller.LoginController).Methods("GET")
	route.HandleFunc("/contact", controller.ContactController)
	route.HandleFunc("/test", controller.DummyController)

	errListener := http.ListenAndServe(":"+port, route)
	if errListener != nil {
		return
	}
}
