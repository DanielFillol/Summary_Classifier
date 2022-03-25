package CSV

import (
	"github.com/Darklabel91/Summary_Classifier"
	"github.com/Darklabel91/Summary_Classifier/Struct"
)

func CreateCSVs(raw []Struct.Raw_decision, nameResultFolder string) {
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
		emp := Summary_Classifier.SummaryClassifier(raw[i].Summary, raw[i].Identifier, raw[i].Court)
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

	exportCSV("totalInfered", nameResultFolder, totalInfered)

	if len(exOfficioReview) != 0 {
		exportCSV("exOfficioReview", nameResultFolder, exOfficioReview)
	}
	if len(diligence) != 0 {
		exportCSV("diligence", nameResultFolder, diligence)
	}
	if len(affected) != 0 {
		exportCSV("affected", nameResultFolder, affected)
	}
	if len(partial) != 0 {
		exportCSV("partial", nameResultFolder, partial)
	}
	if len(groundless) != 0 {
		exportCSV("groundless", nameResultFolder, groundless)
	}
	if len(hasGround) != 0 {
		exportCSV("hasGround", nameResultFolder, hasGround)
	}
	if len(noInfo) != 0 {
		exportCSV("noInfo", nameResultFolder, noInfo)
	}
	if len(notMap) != 0 {
		exportCSV("notMap", nameResultFolder, notMap)
	}
}
