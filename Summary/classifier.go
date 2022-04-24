package Summary

const (
	ExOfficioDec  = "Reexame Necessário"
	DiligenceDec  = "Convertido em Diligência"
	AffectedDec   = "Prejudicado"
	PartialDec    = "Parcial Provimento"
	GroundlessDec = "Improvimento"
	HasGroundDec  = "Provimento"
	NotMapped     = "Não Mapeado"
)

//ReturnSummaryClass return one of seven possible classifications:
//  ExOfficioDec
//	DiligenceDec
//	AffectedDec
// 	PartialDec
//	GroundlessDec
//	HasGroundDec
//	NotMapped
func ReturnSummaryClass(summary string, char int) (string, error) {
	partSummary, err := getLastChar(summary, char)
	if err != nil {
		return "", err
	}
	if exOfficioReview(partSummary, char) {
		return ExOfficioDec, nil
	}
	if diligence(partSummary, char) {
		return DiligenceDec, nil
	}
	if affected(partSummary, char) {
		return AffectedDec, nil
	}
	if partial(partSummary, char) {
		return PartialDec, nil
	}
	if groundless(partSummary, char) {
		return GroundlessDec, nil
	}
	if hasGround(partSummary, char) {
		return HasGroundDec, nil
	}
	return NotMapped, nil
}

//exOfficioReview return true for decision thar need ex officio review (reexame necesário)
func exOfficioReview(summary string, char int) bool {

	less17 := []string{"reexame necessário", "xame necessário"}
	more17 := []string{"reexame necessário"}

	return matchWords(summary, char, less17, more17)

}

//diligence return true for a decision convert in diligence
func diligence(summary string, char int) bool {
	less17 := []string{"conversão do julgamento em diligência", "convertido em diligência", "diligência"}
	more17 := []string{"acórdão em diligência", "conversão do julgamento em diligência", "conversão do feito em diligência", "convertido em diligência", "decisão em diligência", "julgalmento em diligência", "sentença em diligência"}
	return matchWords(summary, char, less17, more17)
}

//affected return true for an affect decision
func affected(summary string, char int) bool {
	less17 := []string{"dicado", "extinção da punibilidade", "icado", "perda do objeto", "prejudicad", "prejudicado", "prejudicada", "perda superveniente do objeto", "prejudicando o exame"}
	more17 := []string{"extinção da punibilidade", "perda do objeto", "prejudicad", "prejudicado", "prejudicada", "perda superveniente do objeto", "prejudicando o exame"}
	return matchWords(summary, char, less17, more17)
}

//groundless return true for a decision with no grounds
func groundless(summary string, char int) bool {
	less17 := []string{"a-se provimento", "acórdão confirmado", "afastad", "ão conhecid", "ção do decidido", "confirmad", "de nenhum deles se conhece", "decisão confirmada", "degenada", "denega", "denegada", "denegado", "deprovid", "desacolhid", "desacolhida", "desacolhido", "descabid", "descabida", "descabido", "descabidos", "descabimento", "desconheci", "desprovi", "desprovid", "desprovida", "desprovido ", "desprovido", "desprovido", "desprovimento do recurso", "desprovimento", "ega conhecimento", "gado provimento", "impossibilidade", "improcedência mantida", "improcedência", "improcedente", "improvid", "impróvid", "improvida", "improvido", "improvimen", "improvimento", "inaplicabilidade do princípio da fungibilidade recursal", "inadimissível", "inadimitida", "inadimitido", "inadmitido", "indefere-se", "indefere", "indeferi", "indeferid", "indeferimento", "indevid", "julgamento confirmado", "mantém-se", "mantém", "mantid", "mantida", "mantido", "manutenção do decidido", "manutenção", "não autorizado", "não cabe conhecer", "não cabe reconhecer", "não conheceram", "não conhecid", "não conhecida", "não conhecido", "não conhecimento", "não provid", "não provimento", "não se conhece", "não se há de", "não há fundamento para acolhimento", "não", "nega conhecimento", "nega provimento", "nega-se provimento", "nega", "negado provimento", "negam-se", "negaram provimento", "nego provimento", "nego", "rejei", "rejeição ", "rejeição", "rejeitad", "rejeitada", "rejeitam-se", "rejeitado", "rejeito", "repelid", "repelido", "repelidos", "retratação não exercid", "se os embargos", "sem alteração do resultado", "sem alteração dos resultados", "sem efeitos modificativos", "sem efeito modificativo", "sentença confirmada"}
	more17 := []string{"acórdão confirmado", "afastad", "de nenhum deles se conhece", "decisão confirmada", "degenada", "denega", "denegada", "denegado", "deprovid", "desacolhid", "desacolhida", "desacolhido", "descabid", "descabida", "descabido", "descabimento", "desconheci", "desprovid", "desprovida", "desprovido ", "desprovido", "desprovimento", "impossibilidade", "improcedência mantida", "improcedência", "improcedente", "impróvid", "improvida", "improvido", "improvimento", "inaplicabilidade do princípio da fungibilidade recursal", "inadimissível", "inadimitida", "inadimitido", "indefere-se", "indefere", "indeferid", "indeferimento", "indevid", "julgamento confirmado", "mantém-se", "mantém", "mantida", "mantido", "manutenção do decidido", "manutenção", "não autorizado", "não cabe conhecer", "não conheceram", "não conhecid", "não conhecida", "não conhecido", "não conhecimento", "não provid", "não provimento", "não se conhece", "não se há de", "não há fundamento para acolhimento", "nega conhecimento", "nega provimento", "nega-se provimento", "negado provimento", "negado conhecimento", "negam-se", "negaram provimento", "nego provimento", "rejeição ", "rejeição", "rejeitad", "repelido", "rejeitam-se", "repelidos", "retratação não exercid", "sem alteração do resultado", "sem alteração dos resultados", "sem efeito modificativo", "sem efeitos modificativos", "sentença confirmada"}
	return matchWords(summary, char, less17, more17)

}

//hasGround return true for a decision with grounds
func hasGround(summary string, char int) bool {
	less17 := []string{"absolver", "absolvição", "absolvida", "absolvido", "acolhem-se-os", "acolhi", "acolhid", "acolhida", "acolhido", "acolhimento", "adequado", "adequação", "admissibilidade", "ajustamento dessa decisão que se impõe", "alterad", "anulad", "anulada", "anulado", "anulem-se", "cabe reconhecer", "cassada", "cassado", "concedid", "conhecid", "concessão da ordem", "correção de erro material", "dá-se provimento", "decisão atacada", "decisaão retratada", "deferid", "deferimento", "devid", "entença anulada", "hecer a omissão", "homologa", "homologad", "homologado", "modificação", "modificad", "possibilidade", "procedência", "procedente", "provid", "provimento", "readequa", "readequação", "recebido", "reconhecer a omissão", "reconheci", "reconhecid", "reconhecimento de ofício", "reformad", "reformada", "reformado", "retratada", "revist", "revogad", "rovido", "sentença anulada"}
	more17 := []string{"absolver", "absolvição", "absolvida", "absolvido", "acolhem-se-os", "acolhida", "acolhido", "acolhimento", "adequado", "adequação", "ajustamento dessa decisão que se impõe", "anulada", "anulado", "anulem-se", "anulação", "cabe reconhecer", "cassada", "cassado", "concedid", "concessão da ordem", "concessão da segurança", "conflito conhecido", "correção de erro material", "dá-se provimento", "decisão atacada", "decisaão retratada", "deferid", "deferimento", "homologad", "modificação", "modificad", "possibilidade", "procedência", "procedente", "provid", "provimento", "readequação", "readequa", "recebid", "reconhecer a omissão", "reconheci", "reconhecid", "reconhecimento de ofício", "reformada", "reformado", "retratada", "revist", "revogad", "rovido", "sentença anulada"}
	return matchWords(summary, char, less17, more17)
}

//partial returns true for a decision with partial grounds
func partial(summary string, char int) bool {
	less17 := []string{"adequação parcial", "almente mantid", "arte conhecid", "cial provimento", "dência parcial", "em parte", "ialmente", "lmente", "mente", "nessa parte", "nesta parte", "neste tópico", "parcial", "parcial provimento", "parcialmente modificada", "parcialmente modificado", "parcialmente provid", "parcialmente", "parte acolhid", "parte apreciad", "parte chonhecid", "parte e provid", "parte provid", "parte, portanto", "provido, em parte", "provido em parte", "rcial provimento", "reformada parcialmente", "te provid"}
	more17 := []string{"adequação parcial", "em parte", "nessa parte", "nesta parte", "parcialmente", "conhecido neste tópico", "em parte acolhid", "em parte apreciad", "em parte conhecid", "em parte e provid", "em parte provid", "em parte, portanto", "parcial provimento", "parcial provimento", "parcialmente mantid", "parcialmente modificada", "parcialmente modificado", "paricalmente provid", "parcialmente provid", "parte conhecid", "procedência parcial", "provido, em parte", "provido em parte", "reformada parcialmente"}
	return matchWords(summary, char, less17, more17)
}
