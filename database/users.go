package database

import responseTypes "github.com/cleanupDev/go_auth_server.git/types"

func GetUserData(email, password string) *responseTypes.UserData {
	db := ConnectDB()
	var userData responseTypes.UserData
	db.Where("email = ? AND password = ?", email, password).First(&userData)
	if userData.Email == "" {
		return nil
	}
	return &userData
}
