package cmd

import (
	"flag"
	"os"

	"github.com/redseverity/forcepath/utils/messages"
	"github.com/redseverity/forcepath/utils/terminal"
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
		messages.Error("Invalid or malformed arguments.")
		messages.Exit()
	}

	return args
}
