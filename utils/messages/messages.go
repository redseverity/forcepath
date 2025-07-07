package messages

import (
	"fmt"
	"os"

	"github.com/redseverity/forcepath/utils/text"
)

func Success(msg string) {
	fmt.Print(text.Bold, text.Green)
	fmt.Print("\n{ ", msg, " }\n\n", text.Reset)
}

func Error(msg string) {
	fmt.Fprint(os.Stderr, text.Bold, text.Red)
	fmt.Fprint(os.Stderr, "\n{ ", msg, " }\n\n", text.Reset)
}

func SuccessInputDetail(label string, param string) {
	fmt.Print(text.Bold, text.Green, "[✓] ")
	fmt.Print(text.Cyan, label, " ", text.Reset, param)
}

func ErrorInputDetail(label string, param string) {
	fmt.Fprint(os.Stderr, text.Bold, text.Red, "[x] ")
	fmt.Fprint(os.Stderr, text.Cyan, label, " ", text.Reset, param)
}

func Exit() {
	fmt.Print(text.Bold, text.Red)
	os.Exit(1)
}

// for future uses;
const (
	Indicator = "●"
)
