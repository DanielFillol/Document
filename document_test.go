package Document

import "testing"

func TestValidateCPF(t *testing.T) {
	tests := []struct {
		cpf      string
		expected bool
	}{
		{"abcdefghijk", false},
		{"123", false},
		{"", false},
		{"12345678900", false},
		{"95524361503", true},
		{"955.243.615-03", true},
		{"  955 243 615 03  ", true},
	}

	for _, tt := range tests {
		actual := ValidateCPF(tt.cpf)
		if actual != tt.expected {
			t.Errorf("ValidateCPF(%s): expected %t, but got %t", tt.cpf, tt.expected, actual)
		}
	}
}

func TestValidateCNPJ(t *testing.T) {
	tests := []struct {
		cnpj     string
		expected bool
	}{
		{"abcdefghijklmn", false},
		{"123", false},
		{"", false},
		{"12345678901234", false},
		{"11222333000181", true},
		{"11.222.333/0001-81", true},
		{"  11 222 333 0001 81  ", true},
	}

	for _, tt := range tests {
		actual := ValidateCNPJ(tt.cnpj)
		if actual != tt.expected {
			t.Errorf("ValidateCNPJ(%s): expected %t, but got %t", tt.cnpj, tt.expected, actual)
		}
	}
}
