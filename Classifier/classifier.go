package Classifier

import (
	"github.com/Darklabel91/Summary_Classifier/Functions"
	"github.com/Darklabel91/Summary_Classifier/Struct"
)

func ClassDecision(summary string, identifier string, court string, char int) Struct.Infered_decision {
	var class string
	var text = Functions.SelectLastChars(summary, char)

	if Functions.ExOfficioReview(text, char) {
		class = "Reexame Necessário"
	} else if Functions.Diligence(text, char) {
		class = "Convertido em Diligência"
	} else if Functions.Affected(text, char) {
		class = "Prejudicado"
	} else if Functions.Partial(text, char) {
		class = "Parcial Provimento"
	} else if Functions.Groundless(text, char) {
		class = "Improvimento"
	} else if Functions.HasGround(text, char) {
		class = "Provimento"
	} else if Functions.Number(text) {
		class = Functions.NumberAnalysis(summary, char+26)
	} else if Functions.Vote(text) {
		class = Functions.VoteAnalysis(summary, char+24)
	} else if text == "Ementa pequena" {
		class = "Sem Informação"
	} else {
		class = "Não Mapeado"
	}

	return Struct.Infered_decision{
		Summary:    summary,
		Text:       text,
		Class:      class,
		Identifier: identifier,
		Court:      court,
	}
}
