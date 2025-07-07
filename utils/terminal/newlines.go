package terminal

import (
	"fmt"
	"strings"
)

func NewLines(count int) {
	if count > 0 {
		fmt.Print(strings.Repeat("\n", count))
	}
}
