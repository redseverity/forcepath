package cmd

import (
	"flag"
	"io"
	"os"

	"github.com/redseverity/forcepath/utils/messages"
)

type Args struct {
	URL     string
	Charset string
}

func GetArgs() Args {
	var args Args

	fs := flag.NewFlagSet("app", flag.ContinueOnError)
	fs.SetOutput(io.Discard) // disable automatic messages

	fs.StringVar(&args.URL, "url", "", "target URL")
	fs.StringVar(&args.Charset, "charset", "", "character set")

	if len(os.Args[1:]) == 0 {
		messages.Error("No flags provided.")
		usage()
	}

	err := fs.Parse(os.Args[1:])
	if err != nil {
		messages.Error("Invalid or malformed arguments.")
		usage()
	}

	checkRequired(args.URL, "-url")
	checkRequired(args.Charset, "-charset")

	return args
}

func checkRequired(value string, name string) {
	if value == "" {
		messages.Info("The " + name + " flag cannot be empty.")
	}
}

func usage() {
	messages.Info("Usage: forcepath -url <target_url> -charset <charset>")
	messages.Exit()
}
