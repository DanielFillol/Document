package Document

import (
	"regexp"
	"strconv"
)

// ValidateCNPJ checks if a given string is a valid Brazilian CNPJ number.
func ValidateCNPJ(cnpj string) bool {
	// Remove all non-numeric characters from the input string
	cnpj = regexp.MustCompile("[^0-9]").ReplaceAllString(cnpj, "")

	// The length of a valid CNPJ number is always 14 digits
	if len(cnpj) != 14 {
		return false
	}

	var sum int
	// The weight values used to calculate the first digit of the CNPJ number
	weight := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	// Calculate the first digit of the CNPJ number
	for i := 0; i < 12; i++ {
		// Convert the i-th character of the input string to an integer
		num, err := strconv.Atoi(string(cnpj[i]))
		if err != nil {
			return false
		}
		// Multiply the i-th digit of the input string by the corresponding weight value
		sum += num * weight[i]
	}

	// Calculate the remainder of the division of the sum by 11
	rem := sum % 11
	if rem < 2 {
		// If the remainder is less than 2, the first digit of the CNPJ number should be 0
		if int(cnpj[12]-'0') != 0 {
			return false
		}
	} else {
		// Otherwise, the first digit of the CNPJ number should be equal to 11 minus the remainder
		if int(cnpj[12]-'0') != 11-rem {
			return false
		}
	}

	sum = 0
	// The weight values used to calculate the second digit of the CNPJ number
	weight = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	// Calculate the second digit of the CNPJ number
	for i := 0; i < 13; i++ {
		// Convert the i-th character of the input string to an integer
		num, err := strconv.Atoi(string(cnpj[i]))
		if err != nil {
			return false
		}
		// Multiply the i-th digit of the input string by the corresponding weight value
		sum += num * weight[i]
	}

	// Calculate the remainder of the division of the sum by 11
	rem = sum % 11
	if rem < 2 {
		// If the remainder is less than 2, the second digit of the CNPJ number should be 0
		if int(cnpj[13]-'0') != 0 {
			return false
		}
	} else {
		// Otherwise, the second digit of the CNPJ number should be equal to 11 minus the remainder
		if int(cnpj[13]-'0') != 11-rem {
			return false
		}
	}

	// If the input string passes all the tests, it is a valid CNPJ number
	return true
}
