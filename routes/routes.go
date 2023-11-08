package routes

import (
	"net/http"
)

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	methodNotAllowedWrapper(handleIndex, http.MethodGet)(w, r)
}

func LoginRoute(w http.ResponseWriter, r *http.Request) {
	switch r.Header.Get("LoginType") {
	case "google":
		methodNotAllowedWrapper(handleGoogleLogin, http.MethodPost)(w, r)
	case "normal":
		methodNotAllowedWrapper(handleNormalLogin, http.MethodPost)(w, r)
	default:
		loginTypes := []string{"google", "normal"}
		w.Header().Set("Allow", joinStrings(loginTypes, ", "))
		http.Error(w, "Login type not supported", http.StatusBadRequest)
	}
}

func RegisterRoute(w http.ResponseWriter, r *http.Request) {
	methodNotAllowedWrapper(handleNormalRegister, http.MethodPost)(w, r)
}
