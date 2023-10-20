package routes

import (
	"encoding/json"
	"log"
	"net/http"

	responseTypes "github.com/cleanupDev/go_auth_server.git/types"
)

func RouteIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")

		responseData := responseTypes.NewResponse("Server Running", nil, nil)

		responseJSON, err := json.Marshal(responseData)
		if err != nil {
			http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
		return

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		response, err := json.Marshal(responseTypes.NewResponse("Method not allowed", nil, nil))
		if err != nil {
			http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
			log.Fatal(err)
			return
		}
		w.Write(response)
		return
	}
}

func RouteLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		response, err := json.Marshal(responseTypes.NewResponse("Method not allowed", nil, nil))
		if err != nil {
			http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
			log.Fatal(err)
			return
		}
		w.Write(response)
		return
	}

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

func RouteRegister(w http.ResponseWriter, r *http.Request) {
	// todo: implement
}
