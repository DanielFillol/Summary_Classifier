package SummaryCSV

import (
	"fmt"
	"github.com/Darklabel91/Summary_Classifier/Summary"
)

//SummaryClassifierCSV classify a given sequence of summary's given by a csv
//and returns it on a given folder with the result csv
func SummaryClassifierCSV(rawFilePath string, separator rune, nameResultFolder string) error {
	raw, err := readCsvFile(rawFilePath, separator)
	if err != nil {
		return err
	}

	err = returnCSVClassSummary(raw, nameResultFolder)
	if err != nil {
		return err
	}

	fmt.Println("Files created")
	return nil
}

//returnCSVClassSummary executes Classify function from a []string
func returnCSVClassSummary(raw []string, nameResultFolder string) error {
	var classifiedSummary []Summary.InferredDecision

	for _, summary := range raw {
		dataReturn, err := Summary.Classify(summary)
		if err != nil {
			classifiedSummary = append(classifiedSummary, Summary.InferredDecision{
				Summary: summary,
				Class:   err.Error(),
			})
		}
		classifiedSummary = append(classifiedSummary, dataReturn)
	}

	err := writeCSV("filesOK", nameResultFolder, classifiedSummary)
	if err != nil {
		return err
	}

	return nil
}
