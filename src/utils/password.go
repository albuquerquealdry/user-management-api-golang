package utils

import (
	"sync"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string, wg *sync.WaitGroup) (string, error) {
	defer wg.Done()
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashBytes), nil
}
