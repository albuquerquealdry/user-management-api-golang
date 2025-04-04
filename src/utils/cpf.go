package utils

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func IsValidCPF(cpf string, wg *sync.WaitGroup, result *bool, errChan chan error) {
	defer wg.Done()

	fmt.Println("Chegou no CPF")
	time.Sleep(10 * time.Second)

	if len(cpf) != 11 {
		*result = false
		errChan <- fmt.Errorf("CPF inválido: tamanho incorreto")
		return
	}

	digits := make([]int, 11)
	for i := 0; i < 11; i++ {
		num, err := strconv.Atoi(string(cpf[i]))
		if err != nil {
			*result = false
			errChan <- fmt.Errorf("CPF inválido: contém caracteres não numéricos")
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

	*result = firstDigit == digits[9] && secondDigit == digits[10]
	if !*result {
		errChan <- fmt.Errorf("CPF inválido: dígitos verificadores incorretos")
	}
}
