package Document

import (
	"regexp"
	"strconv"
)

// validateFirstDigit Calculate the first verification digit of the CPF number by using the first 9 digits of the CPF number and a weighted sum
func validateFirstDigit(cpf string) bool {
	var sum int
	for i := 0; i < 9; i++ {
		// Convert the current digit to an integer
		num, err := strconv.Atoi(string(cpf[i]))
		if err != nil {
			// If the conversion fails, return false
			return false
		}
		// Add the weighted digit to the running sum
		sum += num * (10 - i)
	}

	// Calculate the remainder of the sum divided by 11
	rem := (sum * 10) % 11
	// If the remainder is 10, the verification digit is 0
	if rem == 10 {
		rem = 0
	}

	//Get the second last digit of cpf number
	firstDigit := cpf[len(cpf)-2 : len(cpf)-1]

	// If the calculated verification digit does not match the
	// 10th digit of the CPF number, return false
	if strconv.Itoa(rem) != firstDigit {
		return false
	}

	return true

}

// validateSecondDigit Calculate the second verification digit of the CPF number by using all 10 digits of the CPF number (including the first verification digit that is validated by validateFirstDigit) and a different weighted sum
func validateSecondDigit(cpf string) bool {
	var sum int
	for i := 0; i < 10; i++ {
		// Convert the current digit to an integer
		num, err := strconv.Atoi(string(cpf[i]))
		if err != nil {
			// If the conversion fails, return false
			return false
		}
		// Add the weighted digit to the running sum
		sum += num * (11 - i)
	}

	// Calculate the remainder of the sum divided by 11
	rem := (sum * 10) % 11
	// If the remainder is 10, the verification digit is 0
	if rem == 10 {
		rem = 0
	}

	//Get the last digit of cpf number
	secondDigit := cpf[len(cpf)-1:]

	// If the calculated verification digit does not match the
	// 11th digit of the CPF number, return false
	if strconv.Itoa(rem) != secondDigit {
		return false
	}

	return true
}

// ValidateCPF checks if a given string is a valid Brazilian CPF number.
func ValidateCPF(cpf string) bool {
	// Remove all non-numeric characters from the input string
	cpf = regexp.MustCompile("[^0-9]").ReplaceAllString(cpf, "") // removes all non-numeric characters

	if len(cpf) != 11 {
		return false
	}

	if !validateFirstDigit(cpf) {
		return false
	}

	if !validateSecondDigit(cpf) {
		return false
	}

	// If the function has not yet returned false, the CPF number is valid
	return true
}
