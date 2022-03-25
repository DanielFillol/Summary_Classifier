package Functions

import "strings"

func SelectLastChars(summary string, char int) string {
	var finalText string
	if len(summary) >= char {
		finalText = strings.ToLower(summary[len(summary)-char:])
	} else {
		finalText = "Ementa pequena"
	}
	return finalText
}
