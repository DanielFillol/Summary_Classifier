package Functions

import "strings"

func Number(text string) bool {
	letters := [28]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "x", "w", "y", "z", "/", ")"}
	numbers := [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

	var flag bool
	for _, num := range numbers {
		if strings.Contains(text, num) {
			flag = true
		}

		if flag == true {
			for _, let := range letters {
				if strings.Contains(text, "eros") {
					return true
				} else {
					if strings.Contains(text, let) {
						return false
					}
				}
			}
		}

	}

	return false

}

func NumberAnalysis(summary string, char int) string {
	text := SelectLastChars(summary, len(summary)*2)
	if len(summary) < char {
		return "Sem Informação"
	} else {
		if ExOfficioReview(text, char-26) {
			return "Reexame Necessário"
		} else if Diligence(text, char-26) {
			return "Convertido em Diligência"
		} else if Affected(text, char-26) {
			return "Prejudicado"
		} else if Partial(text, char-26) {
			return "Parcial Provimento"
		} else if Groundless(text, char-26) {
			return "Improvimento"
		} else if HasGround(text, char-26) {
			return "Provimento"
		} else {
			return "Não Mapeado"
		}
	}
}
