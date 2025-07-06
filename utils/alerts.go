package utils

import (
	"fmt"
)

func boldAndRed(text string) {
	ToolName()
	fmt.Print("\n" + BoldText + RedText + text + DefaultText + "\n\n")
}
func boldAndGreen(text string) {
	ToolName()
	fmt.Print("\n" + BoldText + GreenText + text + DefaultText + "\n\n")
}

// green

func ShowEnterToolInfoAlert() {
	boldAndGreen("{ Enter the tool information. }")
}

func ShowChooseOptionAlert() {
	boldAndGreen("{ Choose one of the following options. }")
}

// red

func ShowInvalidOption() {
	boldAndRed("{ Invalid option. Please enter a valid option. }")
}

func ShowInvalidURLAlert() {
	boldAndRed("{ Invalid URL format. Please enter a valid URL. }")
}

func ShowHostUnreachableAlert() {
	boldAndRed("{ Unable to connect. Host is unreachable or does not exist. }")
}

func ShowToolClosedAlert() {
	boldAndRed("{ Tool has closed. }")
}
