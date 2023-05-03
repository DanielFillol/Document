# Brazilian Document Validation

This Go package provides functions to validate Brazilian CPF and CNPJ numbers. CPF (Cadastro de Pessoas Físicas) is a unique identification number given to Brazilian citizens, while CNPJ (Cadastro Nacional da Pessoa Jurídica) is a unique identification number given to Brazilian companies.

## Installation
To use this package, you need to have Go installed on your computer. Then, open a terminal and run the following command:
```go
go get github.com/DanielFillol/Document
```

## Usage
Import the package into your code:
```go
import "github.com/DanielFillol/Document"
```

### ValidateCPF
The ValidateCPF function takes a string as input and returns a boolean value indicating whether the input is a valid CPF number. If the input is not a valid CPF number, the function returns false.
```go
isValid := Document.ValidateCPF("123.456.789-09")
```

### ValidateCNPJ
The ValidateCNPJ function takes a string as input and returns a boolean value indicating whether the input is a valid CNPJ number. If the input is not a valid CNPJ number, the function returns false.
```go
isValid := Document.ValidateCNPJ("12.345.678/0001-00")
```

## How it works
The validation process for both CPF and CNPJ numbers is based on a set of rules defined by the Brazilian government. The rules include checks for the number of digits, the format of the number, and the validity of the verification digits.

The verification digits are calculated using a weighted sum of the digits in the number, where each digit is multiplied by a specific weight. The resulting sum is then divided by a specific value, and the remainder is used to calculate the verification digit.

## Contributing
If you find a bug or have a suggestion for improvement, please open an issue or submit a pull request on GitHub. We welcome contributions from the community.
