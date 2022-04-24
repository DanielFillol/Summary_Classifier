package CSV

import (
	"fmt"
	"github.com/Darklabel91/Summary_Classifier/Summary"
)

//SummaryClassifierCSV classify a given sequence of summary's given on a csv
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
		dataReturn, _ := Summary.Classify(summary)
		classifiedSummary = append(classifiedSummary, dataReturn)
	}

	err := writeCSV("filesOK", nameResultFolder, classifiedSummary)
	if err != nil {
		return err
	}

	return nil
}
