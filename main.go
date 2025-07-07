package main

import (
	"github.com/redseverity/forcepath/cmd"
	"github.com/redseverity/forcepath/validation"
)

func main() {
	args := cmd.GetArgs()

	parsedURL := validation.NormalizeURL(args.URL)

	if parsedURL == "" {
	}
}
