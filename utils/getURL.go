package utils

import (
	"fmt"
	"strings"

	"github.com/redseverity/akfindurl/config"
)

func GetURL() {
	var input string

	fmt.Print(BoldText + RedText + "[+]" + CyanText + " Place URL: " + DefaultText)
	fmt.Scanln(&input)
	input = strings.TrimSpace(input)

	// check if it exists in the Supported_URL_schemes list
	hasScheme := false
	for _, scheme := range config.Supported_URL_schemes {
		if strings.HasPrefix(input, scheme) {
			hasScheme = true
			break
		}
	}

	if !hasScheme {
		input = config.Default_URL_scheme + input
		fmt.Print(clearLine + BoldText + RedText + "[✓]" + CyanText + " Place URL: " + DefaultText + input)
	}

	config.URL = input
}
