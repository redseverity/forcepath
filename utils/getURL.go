package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/redseverity/akfindurl/config"
)

func GetURL() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(BoldText, RedText, "[+]", BlueText, " Place URL: ", DefaultText)
	input, _ := reader.ReadString('\n')
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
		fmt.Print(clearLine, BoldText, RedText, "[+]", BlueText, " Place URL: ", DefaultText, input)
	}

	return input
}
