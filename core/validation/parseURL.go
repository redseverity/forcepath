package validation

import (
	neturl "net/url"
	"regexp"
	"strings"

	"github.com/redseverity/forcepath/core/network"
	"github.com/redseverity/forcepath/utils/messages"
)

func ParseURL(url string, timeout int) string {
	var schemeRegex = regexp.MustCompile(`^https?://`)

	if !schemeRegex.MatchString(url) {
		url = "https://" + url
	}

	if !strings.HasSuffix(url, "/") {
		url += "/"
	}

	parsed, err := neturl.ParseRequestURI(url)
	if err != nil || parsed.Host == "" {
		messages.Error("The URL is malformed. Please enter a valid URL.")
		messages.ErrorInputDetail("target URL:", url)
		messages.Exit1()
	}

	if !network.Request(url, timeout).Verified {
		messages.Error("Unable to connect. Host is unreachable or does not exist.")
		messages.ErrorInputDetail("target URL:", url)
		messages.Exit1()
	}

	return url
}
