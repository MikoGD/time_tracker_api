package utils

import (
	"fmt"
	"regexp"
)

func EscapeString(str string) string {
	escapedString := ""

	pattern := regexp.MustCompile(`[^a-zA-Z0-9\s]`)

	for _, char := range str {
		match := pattern.Find([]byte(string(char)))

		if len(match) > 0 {
			escapedString = escapedString + fmt.Sprintf(`\%s`, string(char))
		} else {
			escapedString = escapedString + string(char)
		}
	}

	return escapedString
}