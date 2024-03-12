package address

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TypeOfAddress(address string) string {
	validTypes := []string{"Rua", "Avenida", "Estrada", "Rodovia"}
	firstWordOfAddress := strings.Split(strings.ToLower(address), " ")[0]

	addressHaveAValidType := false

	for _, address := range validTypes {
		if strings.ToLower(address) == firstWordOfAddress {
			addressHaveAValidType = true
		}
	}
	if addressHaveAValidType {
		converter := cases.Title(language.BrazilianPortuguese)
		return converter.String(firstWordOfAddress)
	}

	return "Invalid Type of Address"
}
