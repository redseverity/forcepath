package cmd

import (
	"flag"
	"fmt"
	neturl "net/url"
	"os"
	"regexp"
	"strings"

	"github.com/redseverity/gosubfinder/global"
	"github.com/redseverity/gosubfinder/network"
	"github.com/redseverity/gosubfinder/utils"
	"github.com/redseverity/gosubfinder/utils/messages"
)

func GetURL() {
	var schemeRegex = regexp.MustCompile(`^https?://`)

	fs := flag.NewFlagSet("url", flag.ContinueOnError)
	url := fs.String("url", "", "target URL")

	err := fs.Parse(os.Args[1:])
	if err != nil {
		utils.ToolName()
		messages.ShowError("{ Invalid or malformed arguments. }")
		messages.ShowExit()
	}

	if *url == "" {
		utils.ToolName()
		messages.ShowError("{ The -url flag is required. }")
		messages.ShowExit()
	}

	// add scheme and trailing slash if missing
	{
		if !schemeRegex.MatchString(*url) {
			*url = "https://" + *url
		}

		if !strings.HasSuffix(*url, "/") {
			*url += "/"
		}
	}

	// validate final URL structure
	parsed, err := neturl.ParseRequestURI(*url)
	if err != nil || parsed.Host == "" {
		utils.ToolName()
		messages.ShowError("{ The URL is malformed. Please enter a valid URL. }")
		fmt.Print(utils.IndicatorFailure, utils.PromptTargetURL, *url, "\n\n")
		messages.ShowExit()
	}

	// check if the URL is reachable
	if !network.CheckURL(*url, 5).Verified {
		utils.ToolName()
		messages.ShowError("{ Unable to connect. Host is unreachable or does not exist. }")
		fmt.Print(utils.IndicatorFailure, utils.PromptTargetURL, *url, "\n\n")
		messages.ShowExit()
	}

	// URL is valid and ready to be used
	utils.ToolName()
	fmt.Fprint(os.Stderr, utils.BoldText, utils.GreenText)
	fmt.Fprint(os.Stderr, "\n{ Parameters configuration. }\n\n")
	fmt.Print(utils.IndicatorSuccess, utils.PromptTargetURL, *url)

	// set the global URL variable
	global.URL = *url
}
