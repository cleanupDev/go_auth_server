package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	responseTypes "github.com/cleanupDev/go_auth_server.git/types"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(user *responseTypes.UserData) (*string, error) {
	// TODO: seperate file read to impede file io errors
	privateKeyBytes, err := os.ReadFile("private.pem")
	if err != nil {
		fmt.Println("Error reading private key:", err)
		return nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		fmt.Println("Error parsing private key:", err)
		return nil, err
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Minute * 60).Unix(),
		"aud":      "http://localhost:5432",
		"iss":      "http://localhost:8080",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		fmt.Println("Error signing the token:", err)
		return nil, err
	}

	return &signedToken, nil

}

func VerifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Header["Authorization"] != nil {
			token, err := jwt.Parse(request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodECDSA)
				if !ok {
					writer.WriteHeader(http.StatusUnauthorized)
					_, err := writer.Write([]byte("You're Unauthorized!"))
					if err != nil {
						return nil, err

					}
				}
				return "", nil

			})
			if err != nil {
				writer.WriteHeader(http.StatusUnauthorized)
				_, err2 := writer.Write([]byte("You're Unauthorized due to error parsing the JWT"))
				if err2 != nil {
					return
				}
			}
			if token.Valid {
				endpointHandler(writer, request)
			} else {
				writer.WriteHeader(http.StatusUnauthorized)
				_, err := writer.Write([]byte("You're Unauthorized due to invalid token"))
				if err != nil {
					return
				}
			}

		} else {
			writer.WriteHeader(http.StatusUnauthorized)
			_, err := writer.Write([]byte("You're Unauthorized due to No token in the header"))
			if err != nil {
				return
			}
		}

	})
}
