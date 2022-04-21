package Functions

import "strings"

func matchWords(text string, char int, less17 []string, more17 []string) bool {
	var listMachtWords []string

	if char < 17 {
		listMachtWords = less17
	} else {
		listMachtWords = more17
	}

	for _, match := range listMachtWords {
		if strings.Contains(text, match) {
			return true
		}
	}

	return false
}
