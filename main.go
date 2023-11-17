package main

import (
	"log"
	"net/http"

	"github.com/cleanupDev/go_auth_server.git/database"
	"github.com/cleanupDev/go_auth_server.git/routes"
	"github.com/cleanupDev/go_auth_server.git/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.SetupDB()
	http.HandleFunc("/", routes.IndexRoute)

	http.HandleFunc("/login", routes.LoginRoute)

	http.HandleFunc("/register", routes.RegisterRoute)

	http.ListenAndServe(server.NewServerConfig("127.0.0.1", "5001").GetAddress(), nil)

}
