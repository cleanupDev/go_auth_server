package routes

import (
	"fmt"
	"io"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get jwt
	tokenString, err := io.ReadAll(r.Body)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to read request body: %s", err.Error())
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	}

	// decode jwt
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(string(tokenString[:]), claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("<YOUR VERIFICATION KEY>"), nil
	})

	htmlString := `
	<html>
		<head>
			<title>Google Login</title>
		</head>
		<body>
			<h1>Google Login</h1>
			<p>JWT: %s</p>
			<p>Claims: %s</p>
		</body>
	</html>
	`

	if err != nil {
		errorMessage := fmt.Sprintf("Failed to parse JWT: %s", err.Error())
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	}

	if !token.Valid {
		errorMessage := fmt.Sprintf("Invalid JWT: %s", err.Error())
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(htmlString, tokenString, claims)))
}
