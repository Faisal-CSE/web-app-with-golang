package main

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	v1API "github.com/web-app-with-golang/api/v1"
	"github.com/web-app-with-golang/controller"
	"github.com/web-app-with-golang/database"
	ptm "github.com/web-app-with-golang/periodic_task_manager"
	env "github.com/web-app-with-golang/project_env"
	"log"
	"net/http"
)

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./favicon.ico")
}

func main() {
	var strAPIVersion1 string

	//DB CONNECTION
	pgConnection := database.PostgresDbConnection()
	if pgConnection == nil {
		log.Fatal("Database connection failed!")
		return
	}

	defer func(pgConnection *sqlx.DB) {
		err := pgConnection.Close()
		if err != nil {
			log.Println("DB connection closed error!")
		}
	}(pgConnection)

	port := env.GoDotEnvVariable("PORT")
	if port == "" {
		port = "8080"
	}

	// TODO SET A PERIODIC TASK FOR EXAMPLE
	ptm.ExecuteACronJob()

	route := mux.NewRouter()

	//Handle fav icon
	route.HandleFunc("/favicon.ico", faviconHandler)

	//Handle static files
	fs := http.FileServer(http.Dir("./assets/"))
	route.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	route.HandleFunc("/", controller.IndexController).Methods(http.MethodGet)
	route.HandleFunc("/register", controller.RegistrationController)
	route.HandleFunc("/login", controller.LoginController)
	route.HandleFunc("/pass-forgot", controller.PasswordForgotController)
	route.HandleFunc("/contact", controller.ContactController)
	route.HandleFunc("/test", controller.DummyController)

	// API URL VERSION 1
	strAPIVersion1 = "/api/v1/"
	route.HandleFunc(strAPIVersion1+"sample-health-check", v1API.SampleHealthCheckAPI).Methods(http.MethodGet)

	log.Print("Starting server on :" + port)
	_ = http.ListenAndServe(":"+port, route)
}
