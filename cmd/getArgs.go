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
	Min     int
	Max     int
}

func GetArgs() Args {
	var args Args

	fs := flag.NewFlagSet("app", flag.ContinueOnError)
	fs.SetOutput(io.Discard) // disable automatic messages

	fs.StringVar(&args.URL, "url", "", "target URL")
	fs.StringVar(&args.Charset, "charset", "abc123", "character set")
	fs.IntVar(&args.Min, "min", 1, "minimum length of generated strings")
	fs.IntVar(&args.Max, "max", 3, "maximum length of generated strings")

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
		messages.Error("The " + name + " flag cannot be empty.")
		messages.Exit()
	}
}

func usage() {
	messages.Info("Usage: forcepath -url <target_url> -charset <charset> [flags]")
	messages.Exit()
}
