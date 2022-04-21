package Functions

import "strings"

func SelectLastChars(summary string, char int) string {
	if len(summary) >= char {
		return strings.ToLower(summary[len(summary)-char:])
	} else {
		return "Ementa pequena"
	}
}
