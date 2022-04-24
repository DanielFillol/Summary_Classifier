package Summary

import "strings"

//matchWords is a simple method that compares a given string on one of possible two arrays
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
