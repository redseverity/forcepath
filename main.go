package main

import (
	"github.com/redseverity/forcepath/cmd"
	"github.com/redseverity/forcepath/core/validation"
	"github.com/redseverity/forcepath/utils"
	"github.com/redseverity/forcepath/utils/messages"
)

func main() {
	utils.ShowBanner()

	args := cmd.GetArgs()

	{
		args.URL = validation.ParseURL(args.URL)
		args.Charset = validation.ParseCharset(args.Charset)
	}

	{
		messages.Success("Parameters loaded.")
		messages.SuccessInputDetail("target URL:", args.URL)
		messages.SuccessInputDetail("charset:", args.Charset)
	}

}
