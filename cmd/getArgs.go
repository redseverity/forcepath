package cmd

import (
	"flag"
	"io"
	"os"
	"strconv"

	"github.com/redseverity/forcepath/utils/messages"
)

type Args struct {
	Help    bool
	URL     string
	Charset string
	Min     int
	Max     int
	Timeout int
	Delay   int
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
	fs.IntVar(&args.Delay, "delay", 0, "delay in milliseconds between each request")

	err := fs.Parse(os.Args[1:])

	if len(os.Args[1:]) == 0 || args.Help {
		messages.Help()
		messages.Exit0()
	}

	if err != nil {
		messages.Error("Invalid or malformed arguments.")
		usage()
	}

	checkEmpty(args.URL, "url")
	checkEmpty(args.Charset, "charset")
	checkMinimum(args.Min, "min", 1)
	checkMinimum(args.Max, "max", 1)
	checkMinimum(args.Timeout, "timeout", 1)
	checkMinimum(args.Delay, "delay", 0)

	return args
}

func checkEmpty(value string, name string) {
	if value == "" {
		messages.Error("The -" + name + " flag cannot be empty.")
		messages.Exit1()
	}
}

func checkMinimum(value int, name string, min int) {
	if value < min {
		messages.Error("The -" + name + " flag must be " + strconv.Itoa(min) + " or greater.")
		messages.Exit1()
	}
}

func usage() {
	messages.Warning("Basic usage: forcepath -url <target_url> -charset <charset> [flags]")
	messages.Exit1()
}
