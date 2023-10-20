package responseTypes

import "fmt"

type UserData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserData(args ...interface{}) (*UserData, error) {
	userData := &UserData{}
	if len(args)%2 != 0 {
		return nil, fmt.Errorf("arguments must be provided in pairs (field name and value)")
	}

	for i := 0; i < len(args); i += 2 {
		fieldName, ok := args[i].(string)
		if !ok {
			return nil, fmt.Errorf("argument at index %d must be a field name (string)", i)
		}
		fieldValue := args[i+1]

		switch fieldName {
		case "Username":
			userData.Username, ok = fieldValue.(string)
		case "Email":
			userData.Email, ok = fieldValue.(string)
		case "Password":
			userData.Password, ok = fieldValue.(string)
		default:
			return nil, fmt.Errorf("unknown field name: %s", fieldName)
		}

		if !ok {
			return nil, fmt.Errorf("value for field %s is of the wrong type", fieldName)
		}
	}

	return userData, nil
}
