package Functions

import "strings"

func Number(text string) bool {
	var ret = false
	var letters = [28]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "x", "w", "y", "z", "/", ")"}
	var numbers = [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	for i := 0; i < len(numbers); i++ {
		if strings.Contains(text, numbers[i]) {
			ret = true
		}
	}
	if ret {
		for a := 0; a < len(letters); a++ {
			if strings.Contains(text, "eros") {
				ret = true
			} else {
				if strings.Contains(text, letters[a]) {
					ret = false
				}
			}
		}
	}
	return ret
}

func NumberAnalysis(summary string, char int) string {
	var ret string
	length := len(summary)
	text := SelectLastChars(summary, length*2)
	if length < char {
		ret = "Sem Informação"
	} else {
		if ExOfficioReview(text, char-26) {
			ret = "Reexame Necessário"
		} else if Diligence(text, char-26) {
			ret = "Convertido em Diligência"
		} else if Affected(text, char-26) {
			ret = "Prejudicado"
		} else if Partial(text, char-26) {
			ret = "Parcial Provimento"
		} else if Groundless(text, char-26) {
			ret = "Improvimento"
		} else if HasGround(text, char-26) {
			ret = "Provimento"
		} else {
			ret = "Não Mapeado"
		}
	}
	return ret
}
