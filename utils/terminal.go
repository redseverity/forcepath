package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/redseverity/forcepath/utils/text"
)

func ShowBanner() {
	Clear()
	fmt.Println(text.Red + "\n" + `███████╗ ██████╗ ██████╗  ██████╗███████╗██████╗  █████╗ ████████╗██╗  ██╗
██╔════╝██╔═══██╗██╔══██╗██╔════╝██╔════╝██╔══██╗██╔══██╗╚══██╔══╝██║  ██║
█████╗  ██║   ██║██████╔╝██║     █████╗  ██████╔╝███████║   ██║   ███████║
██╔══╝  ██║   ██║██╔══██╗██║     ██╔══╝  ██╔═══╝ ██╔══██║   ██║   ██╔══██║
██║     ╚██████╔╝██║  ██║╚██████╗███████╗██║     ██║  ██║   ██║   ██║  ██║
╚═╝      ╚═════╝ ╚═╝  ╚═╝ ╚═════╝╚══════╝╚═╝     ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝` + "\n" + text.Reset)
}

func Clear() {
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

const (
	ProgressBarPrefix = "[00:00:00] "
	ProgressBarFilled = "█"
	ProgressBarEmpty  = "░"
)
