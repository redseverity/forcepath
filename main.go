package main

import (
	"github.com/redseverity/gosubfinder/input"
)

func main() {
	switch input.MainMenu() {
	case "start":
		input.GetURL()

	case "settings":
		break // next code
	}
}
