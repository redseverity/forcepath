package utils

import (
	"fmt"
	"strings"
)

func ToolName() {
	ClearCmd()
	fmt.Println(strings.Repeat(BoldText+RedText+"="+DefaultText, 21))
	fmt.Println(RedText + `        ▌ ▐▘▘   ▌    
▛▌▛▌▛▘▌▌▛▌▜▘▌▛▌▛▌█▌▛▘
▙▌▙▌▄▌▙▌▙▌▐ ▌▌▌▙▌▙▖▌ 
▄▌` + DefaultText)
	fmt.Println(strings.Repeat(BoldText+RedText+"="+DefaultText, 21))
}

const Counter = "[00:00:00] █████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░"
