package login

import (
	"errors"
)

func Login(username string, password string) (string, error) {
	if username == "" || password == "" {
		return "", errors.New("username or password cannot be empty")
	}

	if username == "admin" && password == "password" {
		return "ok", nil
	}

	return "", errors.New("invalid credentials")
}
