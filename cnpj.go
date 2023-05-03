package Document

import (
	"regexp"
	"strconv"
)

// ValidateCNPJ checks if a given string is a valid Brazilian CNPJ number.
func ValidateCNPJ(cnpj string) bool {
	cnpj = regexp.MustCompile("[^0-9]").ReplaceAllString(cnpj, "") // removes all non-numeric characters

	if len(cnpj) != 14 {
		return false
	}

	var sum int
	weight := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	for i := 0; i < 12; i++ {
		num, err := strconv.Atoi(string(cnpj[i]))
		if err != nil {
			return false
		}
		sum += num * weight[i]
	}

	rem := sum % 11
	if rem < 2 {
		if int(cnpj[12]-'0') != 0 {
			return false
		}
	} else {
		if int(cnpj[12]-'0') != 11-rem {
			return false
		}
	}

	sum = 0
	weight = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	for i := 0; i < 13; i++ {
		num, err := strconv.Atoi(string(cnpj[i]))
		if err != nil {
			return false
		}
		sum += num * weight[i]
	}

	rem = sum % 11
	if rem < 2 {
		if int(cnpj[13]-'0') != 0 {
			return false
		}
	} else {
		if int(cnpj[13]-'0') != 11-rem {
			return false
		}
	}

	return true
}
