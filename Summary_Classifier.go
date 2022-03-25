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
	createCSVs(raw, nameResultFolder)
	fmt.Println("Files created")
}

func createCSVs(raw []Struct.Raw_decision, nameResultFolder string) {
	var totalInfered []Struct.Infered_decision
	var exOfficioReview []Struct.Infered_decision
	var diligence []Struct.Infered_decision
	var affected []Struct.Infered_decision
	var partial []Struct.Infered_decision
	var groundless []Struct.Infered_decision
	var hasGround []Struct.Infered_decision
	var noInfo []Struct.Infered_decision
	var notMap []Struct.Infered_decision

	for i := 0; i < len(raw); i++ {
		emp := SummaryClassifier(raw[i].Summary, raw[i].Identifier, raw[i].Court)
		totalInfered = append(totalInfered, emp)
		if emp.Class == "Reexame Necessário" {
			exOfficioReview = append(exOfficioReview, emp)
		} else if emp.Class == "Convertido em Diligência" {
			diligence = append(diligence, emp)
		} else if emp.Class == "Prejudicado" {
			affected = append(affected, emp)
		} else if emp.Class == "Parcial Provimento" {
			partial = append(partial, emp)
		} else if emp.Class == "Improvimento" {
			groundless = append(groundless, emp)
		} else if emp.Class == "Provimento" {
			hasGround = append(hasGround, emp)
		} else if emp.Class == "Sem Informação" {
			noInfo = append(noInfo, emp)
		} else if emp.Class == "Não Mapeado" {
			notMap = append(notMap, emp)
		}

	}

	CSV.ExportCSV("totalInfered", nameResultFolder, totalInfered)

	if len(exOfficioReview) != 0 {
		CSV.ExportCSV("exOfficioReview", nameResultFolder, exOfficioReview)
	}
	if len(diligence) != 0 {
		CSV.ExportCSV("diligence", nameResultFolder, diligence)
	}
	if len(affected) != 0 {
		CSV.ExportCSV("affected", nameResultFolder, affected)
	}
	if len(partial) != 0 {
		CSV.ExportCSV("partial", nameResultFolder, partial)
	}
	if len(groundless) != 0 {
		CSV.ExportCSV("groundless", nameResultFolder, groundless)
	}
	if len(hasGround) != 0 {
		CSV.ExportCSV("hasGround", nameResultFolder, hasGround)
	}
	if len(noInfo) != 0 {
		CSV.ExportCSV("noInfo", nameResultFolder, noInfo)
	}
	if len(notMap) != 0 {
		CSV.ExportCSV("notMap", nameResultFolder, notMap)
	}
}
