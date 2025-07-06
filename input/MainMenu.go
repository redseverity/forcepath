package input

import (
	"fmt"

	"github.com/redseverity/gosubfinder/utils"
)

func MainMenu() string {

	utils.ShowChooseOptionAlert()

	for {
		fmt.Print(utils.BoldText + utils.GreenText)
		fmt.Println("1 - Start the tool")
		fmt.Println("2 - Settings")
		fmt.Print("\n" + utils.IndicatorPending + utils.PromptOption)

		var input string
		fmt.Scanln(&input)

		switch input {
		case "1":
			return "start"
		case "2":
			return "settings"
		}

		utils.ShowInvalidOption()
	}
}
