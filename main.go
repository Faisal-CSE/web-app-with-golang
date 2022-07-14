package main

import (
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

	mux := http.NewServeMux()

	mux.HandleFunc("/", controller.IndexController)
	mux.HandleFunc("/login", controller.LoginController)
	mux.HandleFunc("/contact", controller.ContactController)

	errListener := http.ListenAndServe(":"+port, mux)
	if errListener != nil {
		return
	}
}
