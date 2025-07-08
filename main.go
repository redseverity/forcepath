package main

import (
	"github.com/redseverity/forcepath/cmd"
	"github.com/redseverity/forcepath/utils/terminal"
	"github.com/redseverity/forcepath/validation"
)

func main() {
	terminal.ShowBanner()

	args := cmd.GetArgs()

	parsedURL := validation.NormalizeURL(args.URL)

	if parsedURL == "" {
	}
}
