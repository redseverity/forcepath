package messages

import (
	"fmt"
	"os"

	"github.com/redseverity/gosubfinder/utils"
)

func ShowSuccess(text string) {
	fmt.Print(utils.BoldText + utils.GreenText)
	fmt.Print("\n{ ", text, " }\n\n", utils.DefaultText)
}

func ShowError(text string) {
	fmt.Fprint(os.Stderr, utils.BoldText, utils.RedText)
	fmt.Fprint(os.Stderr, "\n{ ", text, " }\n\n", utils.DefaultText)
}

func ShowExit() {
	fmt.Print(utils.BoldText, utils.RedText)
	os.Exit(1)
}
