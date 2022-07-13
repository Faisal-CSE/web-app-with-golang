package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("<h1>Hello World!</h1>"))
	if err != nil {
		return
	}
}

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

	mux.HandleFunc("/", indexHandler)
	errListener := http.ListenAndServe(":"+port, mux)
	if errListener != nil {
		return
	}
}
