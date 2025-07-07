package validation

import (
	neturl "net/url"
	"regexp"
	"strings"

	"github.com/redseverity/gosubfinder/network"
	"github.com/redseverity/gosubfinder/utils/messages"
	"github.com/redseverity/gosubfinder/utils/terminal"
)

func NormalizeURL(url string) string {
	var schemeRegex = regexp.MustCompile(`^https?://`)

	if url == "" {
		terminal.ShowBanner()
		messages.Error("The -url flag is required.")
		messages.Exit()
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
		messages.Error("The URL is malformed. Please enter a valid URL.")
		messages.ErrorInputDetail("target URL:", url)
		terminal.NewLines(2)
		messages.Exit()
	}

	// check if the URL is reachable
	if !network.CheckURL(url, 5).Verified {
		terminal.ShowBanner()
		messages.Error("Unable to connect. Host is unreachable or does not exist.")
		messages.ErrorInputDetail("target URL:", url)
		terminal.NewLines(2)
		messages.Exit()
	}

	// URL is valid and ready to be used
	terminal.ShowBanner()
	messages.Success("Parameters configuration.")
	messages.SuccessInputDetail("target URL:", url)

	// set the global URL variable
	return url
}
