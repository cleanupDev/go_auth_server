package responseTypes

type UserData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserData(username, email, password string) *UserData {
	return &UserData{
		Username: username,
		Email:    email,
		Password: password,
	}
}
