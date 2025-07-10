package cmd

import (
	"flag"
	"io"
	"os"

	"github.com/redseverity/forcepath/utils/messages"
)

type Args struct {
	Help    bool
	URL     string
	Charset string
	Min     int
	Max     int
}

func GetArgs() Args {
	var args Args

	fs := flag.NewFlagSet("app", flag.ContinueOnError)
	fs.SetOutput(io.Discard) // disable automatic messages

	fs.BoolVar(&args.Help, "help", false, "show help message")
	fs.StringVar(&args.URL, "url", "", "target URL")
	fs.StringVar(&args.Charset, "charset", "abc123", "character set")
	fs.IntVar(&args.Min, "min", 1, "minimum length of generated strings")
	fs.IntVar(&args.Max, "max", 3, "maximum length of generated strings")

	err := fs.Parse(os.Args[1:])

	if len(os.Args[1:]) == 0 {
		messages.Error("No flags provided.")
		usage()
	}

	if args.Help {
		messages.Help()
		os.Exit(0)
	}

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
	messages.Warning("Usage: forcepath -url <target_url> -charset <charset> [flags]")
	messages.Exit()
}
