package validation

import (
	"fmt"
	neturl "net/url"
	"regexp"
	"strings"

	"github.com/redseverity/gosubfinder/network"
	"github.com/redseverity/gosubfinder/utils"
	"github.com/redseverity/gosubfinder/utils/messages"
	"github.com/redseverity/gosubfinder/utils/terminal"
)

func NormalizeURL(url string) string {
	var schemeRegex = regexp.MustCompile(`^https?://`)

	if url == "" {
		terminal.ShowBanner()
		messages.ShowError("The -url flag is required.")
		messages.ShowExit()
	}

	// add scheme and trailing slash if missing
	{
		if !schemeRegex.MatchString(url) {
			url = "https://" + url
		}

		if !strings.HasSuffix(url, "/") {
			url += "/"
		}
	}

	// validate final URL structure
	parsed, err := neturl.ParseRequestURI(url)
	if err != nil || parsed.Host == "" {
		terminal.ShowBanner()
		messages.ShowError("The URL is malformed. Please enter a valid URL.")
		fmt.Print(utils.IndicatorFailure, utils.PromptTargetURL, url, "\n\n")
		messages.ShowExit()
	}

	// check if the URL is reachable
	if !network.CheckURL(url, 5).Verified {
		terminal.ShowBanner()
		messages.ShowError("Unable to connect. Host is unreachable or does not exist.")
		fmt.Print(utils.IndicatorFailure, utils.PromptTargetURL, url, "\n\n")
		messages.ShowExit()
	}

	// URL is valid and ready to be used
	terminal.ShowBanner()
	messages.ShowSuccess("Parameters configuration.")
	fmt.Print(utils.IndicatorSuccess, utils.PromptTargetURL, url)

	// set the global URL variable
	return url
}
