package input

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"

	urlsettings "github.com/redseverity/gosubfinder/settings/url"
	"github.com/redseverity/gosubfinder/utils"
)

func GetURL() {
	scanner := bufio.NewScanner(os.Stdin)
	var raw string

	utils.ShowEnterToolInfoAlert()

	for {

		fmt.Print(utils.PrefixPending + utils.PromptTargetURL)

		// read and validate basic inputc
		{
			if !scanner.Scan() {
				utils.ShowToolClosedAlert()
				return
			}

			raw = scanner.Text()
			raw = strings.TrimSpace(raw)

			if raw == "" || strings.Contains(raw, " ") {
				utils.ShowInvalidURLAlert()
				continue
			}

		}

		// add scheme and trailing slash if missing
		{
			if !strings.HasPrefix(raw, "http://") && !strings.HasPrefix(raw, "https://") {
				raw = urlsettings.DefaultScheme + raw
			}

			if !strings.HasSuffix(raw, "/") {
				raw += "/"
			}
		}

		// validate final URL structure
		parsed, err := url.ParseRequestURI(raw)
		if err != nil || parsed.Host == "" {
			utils.ShowInvalidURLAlert()
			continue
		}

		break
	}

	// URL is valid and ready to be used
	utils.ShowEnterToolInfoAlert()
	fmt.Print(utils.PrefixSuccess, utils.PromptTargetURL, raw)

	urlsettings.URL = raw
}
