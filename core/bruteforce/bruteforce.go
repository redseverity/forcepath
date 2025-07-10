package bruteforce

import (
	"fmt"
)

func Run(charset string, min int, max int, url string, timeout int) {
	var generate func(prefix string, rest int)

	generate = func(prefix string, rest int) {
		if rest == 0 {
			fmt.Println(url + prefix) // Output the generated string
			return
		}
		for i := 0; i < len(charset); i++ {
			generate(prefix+string(charset[i]), rest-1)
		}
	}

	for length := min; length <= max; length++ {
		generate("", length)
	}
}
