package Summary_Classifier

import (
	"fmt"
	"github.com/Darklabel91/Summary_Classifier/CSV"
	"github.com/Darklabel91/Summary_Classifier/Classifier"
	"github.com/Darklabel91/Summary_Classifier/Struct"
)

func SummaryClassifier(summary string, identifier string, court string) Struct.Infered_decision {

	ret := Classifier.ClassDecision(summary, identifier, court, 16)

	if ret.Class == "Não Mapeado" {
		ret = Classifier.ClassDecision(summary, identifier, court, 70)
		if ret.Class == "Não Mapeado" || ret.Class == "Sem informaçao" {
			ret = Classifier.ClassDecision(summary, identifier, court, len(summary)/2)
			if ret.Class == "Não Mapeado" {
				ret = Classifier.ClassDecision(summary, identifier, court, len(summary))
			}
		}
	}

	return ret

}
func SummaryClassifierCSV(rawDecisionPath string, separator rune, nameResultFolder string) {
	raw := CSV.ReadCsvFile(rawDecisionPath, separator)
	CSV.CreateCSVs(raw, nameResultFolder)
	fmt.Println("Files created")
}
