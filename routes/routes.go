package routes

import (
	"net/http"
)

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	methodNotAllowedWrapper(indexLogic, http.MethodGet)(w, r)
}

func LoginRoute(w http.ResponseWriter, r *http.Request) {
	methodNotAllowedWrapper(loginLogic, http.MethodPost)(w, r)
}

func RegisterRoute(w http.ResponseWriter, r *http.Request) {
	methodNotAllowedWrapper(registerLogic, http.MethodPost)(w, r)
}
