package utils

import (
	"strconv"
	"sync"
)

func IsValidCPF(cpf string, wg *sync.WaitGroup) bool {
	defer wg.Done()
	if len(cpf) != 11 {
		return false
	}

	digits := make([]int, 11)
	for i := 0; i < 11; i++ {
		num, err := strconv.Atoi(string(cpf[i]))
		if err != nil {
			return false
		}
		digits[i] = num
	}
	sum := 0
	for i := 0; i < 9; i++ {
		sum += digits[i] * (10 - i)
	}
	firstDigit := 11 - (sum % 11)
	if firstDigit >= 10 {
		firstDigit = 0
	}

	sum = 0
	for i := 0; i < 10; i++ {
		sum += digits[i] * (11 - i)
	}
	secondDigit := 11 - (sum % 11)
	if secondDigit >= 10 {
		secondDigit = 0
	}

	return firstDigit == digits[9] && secondDigit == digits[10]
}
