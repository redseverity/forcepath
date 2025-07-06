package main

import (
	"github.com/redseverity/gosubfinder/cmd"
	"github.com/redseverity/gosubfinder/validation"
)

func main() {
	args := cmd.GetArgs()

	parsedURL := validation.NormalizeURL(args.URL)

	if parsedURL == "" {
	}
}
