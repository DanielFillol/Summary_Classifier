package Functions

import "strings"

func Vote(text string) bool {
	var vote = [1]string{"voto n"}

	for _, vt := range vote {
		if strings.Contains(text, vt) {
			return true
		}
	}

	return false
}

func VoteAnalysis(summary string, char int) string {
	text := SelectLastChars(summary, char)
	if ExOfficioReview(text, char-24) {
		return "Reexame Necessário"
	} else if Diligence(text, char-24) {
		return "Convertido em Diligência"
	} else if Affected(text, char-24) {
		return "Prejudicado"
	} else if Partial(text, char-24) {
		return "Parcial Provimento"
	} else if Groundless(text, char-24) {
		return "Improvimento"
	} else if HasGround(text, char-24) {
		return "Provimento"
	} else {
		return "Não Mapeado"
	}
}
