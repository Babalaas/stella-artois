package service

import (
	"errors"
)

func comparePasswords(storedPassword string, suppliedPassword string) (bool, error) {
	if storedPassword == suppliedPassword {
		return true, nil
	}

	return false, errors.New("Password Service: Passwords do not match")
}
