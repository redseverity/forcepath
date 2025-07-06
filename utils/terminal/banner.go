package terminal

import (
	"fmt"
	"strings"

	"github.com/redseverity/gosubfinder/utils"
)

func printVisualDivider() {
	fmt.Println(strings.Repeat(utils.BoldText+utils.RedText+"="+utils.DefaultText, 21))
}

func ShowBanner() {
	Clear()
	printVisualDivider()
	fmt.Println(utils.RedText + `        ▌ ▐▘▘   ▌    
▛▌▛▌▛▘▌▌▛▌▜▘▌▛▌▛▌█▌▛▘
▙▌▙▌▄▌▙▌▙▌▐ ▌▌▌▙▌▙▖▌ 
▄▌` + utils.DefaultText)
	printVisualDivider()
}
