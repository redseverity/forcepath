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
	Timeout int
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
	fs.IntVar(&args.Timeout, "timeout", 3, "timeout in seconds for each request")

	err := fs.Parse(os.Args[1:])

	if len(os.Args[1:]) == 0 {
		messages.Help()
		messages.Exit0()
	}

	if args.Help {
		messages.Help()
		messages.Exit0()
	}

	if err != nil {
		messages.Error("Invalid or malformed arguments.")
		usage()
	}

	checkEmpty(args.URL, "url")
	checkEmpty(args.Charset, "charset")
	checkMinimum(args.Min, "min")
	checkMinimum(args.Max, "max")
	checkMinimum(args.Timeout, "timeout")

	return args
}

func checkEmpty(value string, name string) {
	if value == "" {
		messages.Error("The -" + name + " flag cannot be empty.")
		messages.Exit1()
	}
}

func checkMinimum(value int, name string) {
	if value < 1 {
		messages.Error("The -" + name + " flag must be 1 or greater.")
		messages.Exit1()
	}
}

func usage() {
	messages.Warning("Basic usage: forcepath -url <target_url> -charset <charset> [flags]")
	messages.Exit1()
}
