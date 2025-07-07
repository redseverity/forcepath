package terminal

import (
	"fmt"
	"strings"

	"github.com/redseverity/gosubfinder/utils/text"
)

func printVisualDivider() {
	fmt.Print(text.Bold, text.Red)
	fmt.Println(strings.Repeat("=", 21))
	fmt.Print(text.Reset)
}

func ShowBanner() {
	Clear()
	printVisualDivider()
	fmt.Println(text.Red + `        ▌ ▐▘▘   ▌    
▛▌▛▌▛▘▌▌▛▌▜▘▌▛▌▛▌█▌▛▘
▙▌▙▌▄▌▙▌▙▌▐ ▌▌▌▙▌▙▖▌ 
▄▌` + text.Reset)
	printVisualDivider()
}
