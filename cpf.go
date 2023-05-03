package Document

import (
	"regexp"
	"strconv"
)

// ValidateCPF checks if a given string is a valid Brazilian CPF number.
func ValidateCPF(cpf string) bool {
	cpf = regexp.MustCompile("[^0-9]").ReplaceAllString(cpf, "") // removes all non-numeric characters

	if len(cpf) != 11 {
		return false
	}

	var sum int
	for i := 0; i < 9; i++ {
		num, err := strconv.Atoi(string(cpf[i]))
		if err != nil {
			return false
		}
		sum += num * (10 - i)
	}

	rem := sum % 11
	if rem == 10 {
		rem = 0
	}
	if rem != int(cpf[9]-'0') {
		return false
	}

	sum = 0
	for i := 0; i < 10; i++ {
		num, err := strconv.Atoi(string(cpf[i]))
		if err != nil {
			return false
		}
		sum += num * (11 - i)
	}

	rem = sum % 11
	if rem == 10 {
		rem = 0
	}
	if rem != int(cpf[10]-'0') {
		return false
	}

	return true
}
