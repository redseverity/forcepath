package bruteforce

import (
	"strconv"
	"time"

	"github.com/redseverity/forcepath/core/network"
	"github.com/redseverity/forcepath/utils/messages"
)

func Run(charset string, min int, max int, url string, timeout int, delay int) {
	var generate func(prefix string, rest int)

	generate = func(prefix string, rest int) {

		if delay > 0 {
			time.Sleep(time.Duration(delay) * time.Millisecond)
		}

		if rest == 0 {
			var request = network.Request(url+prefix, timeout)

			if !request.Verified {
				messages.ErrorRequest(strconv.Itoa(request.Status)+":", request.URL)
				return
			}

			messages.SuccessRequest(strconv.Itoa(request.Status)+":", request.URL)

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
