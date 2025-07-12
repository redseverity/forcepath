package main

import (
	"fmt"
	"strconv"

	"github.com/redseverity/forcepath/cmd"
	"github.com/redseverity/forcepath/core/bruteforce"
	"github.com/redseverity/forcepath/core/validation"
	"github.com/redseverity/forcepath/utils"
	"github.com/redseverity/forcepath/utils/messages"
)

func main() {
	utils.ShowBanner()

	args := cmd.GetArgs()

	// Validação dos parâmetros
	args.URL = validation.ParseURL(args.URL, args.Timeout)
	args.Charset = validation.ParseCharset(args.Charset)

	// Exibição dos parâmetros carregados
	messages.Success("Parameters loaded.")
	messages.SuccessInputDetail("Target URL:", args.URL)
	messages.SuccessInputDetail("Charset:", args.Charset)
	messages.SuccessInputDetail("Min length:", strconv.Itoa(args.Min))
	messages.SuccessInputDetail("Max length:", strconv.Itoa(args.Max))
	messages.SuccessInputDetail("Timeout:", strconv.Itoa(args.Timeout)+"s")
	messages.SuccessInputDetail("Delay:", strconv.Itoa(args.Delay)+"ms\n")

	// Execução do brute-force
	messages.Success("Starting brute-force scan...")
	bruteforce.Run(args.Charset, args.Min, args.Max, args.URL, args.Timeout, args.Delay)
	fmt.Println()
	messages.Success("Brute-force scan completed.")
}
