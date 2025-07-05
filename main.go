package main

import (
	"fmt"

	"github.com/redseverity/gosubfinder/utils"
)

func main() {
	utils.ClearCmd()
	utils.ToolName()

	option := utils.OptionsPanel()

	switch option {
	case "start":
		utils.ClearCmd()
		utils.ToolName()
		fmt.Println(utils.BoldText + utils.GreenText + "\n{ Put the tool information }\n" + utils.DefaultText)
		fmt.Print(utils.BoldText + utils.RedText + "[+]" + utils.CyanText + " target URL: " + utils.DefaultText)
		utils.GetURL()

	case "config":
		break // next code
	}
}
