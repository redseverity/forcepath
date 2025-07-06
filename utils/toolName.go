package utils

import (
	"fmt"
	"strings"
)

func printVisualDivider() {
	fmt.Println(strings.Repeat(BoldText+RedText+"="+DefaultText, 21))
}

func ToolName() {
	ClearCmd()
	printVisualDivider()
	fmt.Println(RedText + `        ▌ ▐▘▘   ▌    
▛▌▛▌▛▘▌▌▛▌▜▘▌▛▌▛▌█▌▛▘
▙▌▙▌▄▌▙▌▙▌▐ ▌▌▌▙▌▙▖▌ 
▄▌` + DefaultText)
	printVisualDivider()
}
