package utils

import "fmt"

func ShowExitToolAlert() {
	ToolName()
	fmt.Println(BoldText + RedText + "\n{ Tool has closed. }\n" + DefaultText)
}

func ShowURLFormatAlert() {
	ToolName()
	fmt.Println(BoldText + RedText + "\n{ Invalid URL format, Please try again }\n" + DefaultText)
	fmt.Print(BoldText + RedText + "[+]" + CyanText + " target URL: " + DefaultText)
}
