package utils

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/redseverity/akfindurl/config"
)

func errorAlert() {
	ClearCmd()
	ToolName()
	fmt.Println(BoldText + RedText + "\n{ Invalid URL format, Please try again }\n" + DefaultText)
	fmt.Print(BoldText + RedText + "[+]" + CyanText + " target URL: " + DefaultText)
}

func GetURL() {

	var raw string
	verifyScheme := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9+.-]*://`)

	for {

		fmt.Scanln(&raw)
		raw = strings.TrimSpace(raw)

		if raw == "" {
			errorAlert()
			continue
		}

		if !verifyScheme.MatchString(raw) {
			raw = config.Default_Scheme + raw
		}

		parsed, err := url.ParseRequestURI(raw)
		if err != nil || parsed.Host == "" {
			errorAlert()
			continue
		}

		break
	}

	fmt.Print(clearLine + BoldText + RedText + "[✓]" + CyanText + " target URL: " + DefaultText + raw)
	config.URL = raw
}
