package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/cleanupDev/go_auth_server.git/database"
	"github.com/cleanupDev/go_auth_server.git/handlers"
	responseTypes "github.com/cleanupDev/go_auth_server.git/types"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
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

func handleNormalLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to read request body: %s", err.Error())
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	}

	user, err := responseTypes.NewUserData()
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to parse user data: %s", err.Error())
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
		return
	}

	userDB := database.GetUserData(user.Email, user.Password)
	if userDB == nil {
		http.Error(w, "Credentials incorrect", http.StatusNotFound)
		return
	}

	accessString, err := handlers.GenerateJWT(user)
	if err != nil {
		log.Printf("Failed to generate access key: %s", err.Error())
		http.Error(w, "Failed to generate access key", http.StatusInternalServerError)
		return
	}
	accessKey := responseTypes.NewAccessKey(accessString)

	responseData := responseTypes.NewResponse("Login successful", user, accessKey)

	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func handleNormalRegister(w http.ResponseWriter, r *http.Request) {
	// TODO: implement register logic
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
