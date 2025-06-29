package utils

import (
	"os"
	"os/exec"
	"runtime"
)

const (
	DefaultText = "\033[0m"
	BoldText    = "\033[1m"
	clearLine   = "\033[1A\033[2K\r"

	// colors
	BlackText   = "\033[30m"
	RedText     = "\033[31m"
	GreenText   = "\033[32m"
	YellowText  = "\033[33m"
	BlueText    = "\033[34m"
	MagentaText = "\033[35m"
	CyanText    = "\033[36m"
	WhiteText   = "\033[37m"
)

func ClearCmd() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default: // Linux, macOS, etc
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
