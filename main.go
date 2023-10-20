package main

import (
	"net/http"

	"github.com/cleanupDev/go_auth_server.git/routes"
	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/", routes.IndexRoute)

	http.HandleFunc("/login", routes.LoginRoute)

	http.HandleFunc("/register", routes.RegisterRoute)

	http.ListenAndServe(NewServerConfig("127.0.0.1", "8080").GetAddress(), nil)

}