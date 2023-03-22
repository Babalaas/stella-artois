package service

import (
	"errors"
)

func comparePasswords(storedPassword string, suppliedPassword string) (bool, error) {
	if storedPassword == suppliedPassword {
		return true, nil
	}

	return false, errors.New("password service: passwords do not match")
}
