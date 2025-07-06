package cmd

import (
	"flag"
	"os"

	"github.com/redseverity/gosubfinder/utils/messages"
	"github.com/redseverity/gosubfinder/utils/terminal"
)

type Args struct {
	URL     string
	Charset string
}

func GetArgs() Args {
	var args Args

	fs := flag.NewFlagSet("app", flag.ContinueOnError)

	fs.StringVar(&args.URL, "url", "", "target URL")
	fs.StringVar(&args.Charset, "charset", "", "character set")

	err := fs.Parse(os.Args[1:])

	if err != nil {
		terminal.ShowBanner()
		messages.ShowError("Invalid or malformed arguments.")
		messages.ShowExit()
	}

	return args
}
