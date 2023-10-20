package routes

import (
	"encoding/json"
	"net/http"

	responseTypes "github.com/cleanupDev/go_auth_server.git/types"
)

func indexLogic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	responseData := responseTypes.NewResponse("Server Running", nil, nil)

	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func loginLogic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// TODO: implement login logic

	accessKey := responseTypes.NewAccessKey()
	userData := responseTypes.NewUserData("test_username", "test_email", "test_password")

	responseData := responseTypes.NewResponse("Login successful", userData, accessKey)

	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func registerLogic(w http.ResponseWriter, r *http.Request) {
	// TODO: implement register logic
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
