package utils

import (
	"fmt"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string, wg *sync.WaitGroup, result *string, errChan chan error) {
	defer wg.Done()

	fmt.Println("Chegou no  HASH")
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		errChan <- err
		return
	}
	*result = string(hashBytes)
}
