package validation

import (
	"regexp"
	"strings"

	"github.com/redseverity/forcepath/utils/messages"
)

func ParseCharset(charset string) string {
	var validCharset = regexp.MustCompile(`^[a-zA-Z0-9\-_.~]+$`)

	if !validCharset.MatchString(charset) {
		messages.Error("Invalid charset provided.")
		messages.Warning("Supported: a-z  A-Z  0-9  -  _  .  ~")
		messages.Exit1()
	}

	var seen = make(map[rune]bool)
	var cleaned strings.Builder

	for _, char := range charset {
		if !seen[char] {
			seen[char] = true
			cleaned.WriteRune(char)
		}
	}

	return cleaned.String()
}
