package database

import (
	"log"

	responseTypes "github.com/cleanupDev/go_auth_server.git/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	//! TODO: move this to a config file
	dsn := "user=myuser password=mypassword dbname=mydb host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func SetupDB() {
	db := ConnectDB()

	db.AutoMigrate(&responseTypes.UserData{})
}
