package Functions

import "strings"

func matchWords(text string, char int, less17 []string, more17 []string) bool {
	var listMachtWords []string
	var ret = false

	if char < 17 {
		listMachtWords = less17
	} else {
		listMachtWords = more17
	}

	for i := 0; i < len(listMachtWords); i++ {
		if strings.Contains(text, listMachtWords[i]) {
			ret = true
		}
	}

	return ret
}
