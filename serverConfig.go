package main

import "fmt"

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
