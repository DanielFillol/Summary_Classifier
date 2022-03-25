package Functions

import "strings"

func Vote(text string) bool {
	var ret = false
	var vote = [1]string{"voto n"}
	for i := 0; i < len(vote); i++ {
		if strings.Contains(text, vote[i]) {
			ret = true
		}
	}
	return ret
}

func VoteAnalysis(summary string, char int) string {
	text := SelectLastChars(summary, char)
	var tipo string
	if ExOfficioReview(text, char-24) {
		tipo = "Reexame Necessário"
	} else if Diligence(text, char-24) {
		tipo = "Convertido em Diligência"
	} else if Affected(text, char-24) {
		tipo = "Prejudicado"
	} else if Partial(text, char-24) {
		tipo = "Parcial Provimento"
	} else if Groundless(text, char-24) {
		tipo = "Improvimento"
	} else if HasGround(text, char-24) {
		tipo = "Provimento"
	} else {
		tipo = "Não Mapeado"
	}
	return tipo
}
