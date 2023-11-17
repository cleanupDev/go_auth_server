package server

import (
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/urlshortener/v1"
)

type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func NewServerConfig(host, port string) *ServerConfig {
	return &ServerConfig{
		Host: host,
		Port: port,
	}
}

func (c *ServerConfig) GetAddress() string {
	fmt.Println("Server running on: " + c.Host + ":" + c.Port)
	return c.Host + ":" + c.Port
}

var GoogleOauthConfig = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	Endpoint:     google.Endpoint,
	Scopes:       []string{urlshortener.UrlshortenerScope},
}
