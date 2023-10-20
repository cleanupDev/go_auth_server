package routes

import (
	"net/http"
)

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	methodNotAllowedWrapper(handleIndex, http.MethodGet)(w, r)
}

func LoginRoute(w http.ResponseWriter, r *http.Request) {
	methodNotAllowedWrapper(handleLogin, http.MethodPost)(w, r)
}

func RegisterRoute(w http.ResponseWriter, r *http.Request) {
	methodNotAllowedWrapper(handleRegister, http.MethodPost)(w, r)
}
