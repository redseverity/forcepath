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

	{
		fs.StringVar(&args.URL, "url", "", "target URL")
		fs.StringVar(&args.Charset, "charset", "", "character set")
	}

	err := fs.Parse(os.Args[1:])
	if err != nil {
		messages.Error("Invalid or malformed arguments.")
		usage()
	}

	if args.URL == "" {
		messages.Error("The -url flag cannot be empty.")
		usage()
	}

	if args.Charset == "" {
		messages.Error("The -charset flag cannot be empty.")
		usage()
	}

	return args
}

func usage() {
	messages.Info("Usage: forcepath -url <target_url> -charset <charset>")
	messages.Exit()
}
