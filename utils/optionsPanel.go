package utils

import (
	"fmt"
)

func OptionsPanel() string {

	fmt.Println(BoldText + GreenText + "\n{ Select one of the following options. }\n" + DefaultText)

	for {
		fmt.Println(BoldText + GreenText + "1 - Start the tool")
		fmt.Println("2 - Settings" + DefaultText)
		fmt.Print(BoldText + RedText + "\n[+] " + CyanText + "Option: " + DefaultText)

		var input string
		fmt.Scanln(&input)

		switch input {
		case "1":
			return "start"
		case "2":
			return "config"
		}

		ClearCmd()
		ToolName()
		fmt.Println(BoldText + RedText + "\n{ Invalid input. Please enter a valid option. }\n" + DefaultText)
	}
}
