package cmd

import (
	"flag"
	"fmt"
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

	err := fs.Parse(os.Args[1:])

	if err != nil {
		messages.Error("Invalid or malformed arguments.")
		fmt.Print("Usage: forcepath -url <target_url> -charset <charset>\n\n")
		messages.Exit()
	}

	return args
}
