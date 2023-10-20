package main

import (
	"net/http"

	"github.com/cleanupDev/go_auth_server.git/database"
	"github.com/cleanupDev/go_auth_server.git/routes"
	_ "github.com/lib/pq"
)

func main() {
	database.SetupDB()
	http.HandleFunc("/", routes.IndexRoute)

	http.HandleFunc("/login", routes.LoginRoute)

	http.HandleFunc("/register", routes.RegisterRoute)

	http.ListenAndServe(NewServerConfig("0.0.0.0", "5001").GetAddress(), nil)

}
