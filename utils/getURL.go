package utils

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/redseverity/gosubfinder/config"
)

func GetURL() {
	scanner := bufio.NewScanner(os.Stdin)
	var raw string

	for {
		if !scanner.Scan() {
			ShowExitToolAlert()
			return
		}

		raw = scanner.Text()
		raw = strings.TrimSpace(raw)

		if raw == "" || strings.Contains(raw, " ") {
			ShowURLFormatAlert()
			continue
		}

		if !strings.HasPrefix(raw, "http://") && !strings.HasPrefix(raw, "https://") {
			raw = config.Default_Scheme + raw
		}

		parsed, err := url.ParseRequestURI(raw)
		if err != nil || parsed.Host == "" {
			ShowURLFormatAlert()
			continue
		}

		break
	}

	fmt.Print(clearLine + BoldText + RedText + "[✓]" + CyanText + " target URL: " + DefaultText + raw)
	config.URL = raw
}
