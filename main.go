package main

import (
	"fmt"

	"github.com/redseverity/akfindurl/utils"
)

func main() {
	utils.ClearCmd()
	utils.ToolName()

	option := utils.OptionsPanel()

	switch option {
	case "start":
		utils.ClearCmd()
		utils.ToolName()
		fmt.Println(utils.BoldText + utils.GreenText + "\n{ put the tool information }\n" + utils.DefaultText)
		utils.GetURL()

	case "config":
		break // next code
	}
}
