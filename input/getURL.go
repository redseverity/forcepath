package input

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/redseverity/gosubfinder/config"
	"github.com/redseverity/gosubfinder/utils"
)

func GetURL() {
	scanner := bufio.NewScanner(os.Stdin)
	var raw string

	utils.ShowEnterToolInfoAlert()
	fmt.Print(utils.PrefixPending + utils.PromptTargetURL)

	for {
		if !scanner.Scan() {
			utils.ShowToolClosedAlert()
			return
		}

		raw = scanner.Text()
		raw = strings.TrimSpace(raw)

		if raw == "" || strings.Contains(raw, " ") {
			utils.ShowInvalidURLAlert()
			fmt.Print(utils.PrefixPending + utils.PromptTargetURL)
			continue
		}

		if !strings.HasPrefix(raw, "http://") && !strings.HasPrefix(raw, "https://") {
			raw = config.Default_Scheme + raw
		}

		parsed, err := url.ParseRequestURI(raw)
		if err != nil || parsed.Host == "" {
			utils.ShowInvalidURLAlert()
			fmt.Print(utils.PrefixPending + utils.PromptTargetURL)
			continue
		}

		break
	}

	utils.ShowEnterToolInfoAlert()
	fmt.Print(utils.PrefixSuccess, utils.PromptTargetURL, raw)

	config.URL = raw
}
