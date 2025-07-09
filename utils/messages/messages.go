package messages

import (
	"fmt"
	"os"

	"github.com/redseverity/forcepath/utils/terminal"
	"github.com/redseverity/forcepath/utils/text"
)

func Success(msg string) {
	fmt.Print(text.Bold, text.Green)
	fmt.Print("{ ", msg, " }", text.Reset)
	terminal.NewLines(2)
}

func Error(msg string) {
	fmt.Fprint(os.Stderr, text.Bold, text.Red)
	fmt.Fprint(os.Stderr, "{ ", msg, " }", text.Reset)
	terminal.NewLines(2)
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
