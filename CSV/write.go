package CSV

import (
	"encoding/csv"
	"github.com/Darklabel91/Summary_Classifier/Summary"
	"os"
	"path/filepath"
)

//ExportCSV exports a csv to a given folder, with a given name from a collection of AnalysisCNJ
func writeCSV(fileName string, folderName string, decisions []Summary.InferredDecision) error {
	var rows [][]string

	rows = append(rows, generateHeaders())

	for _, decision := range decisions {
		rows = append(rows, generateRow(decision))
	}

	cf, err := createFile(folderName + "/" + fileName + ".csv")
	if err != nil {
		return err
	}

	defer cf.Close()

	w := csv.NewWriter(cf)

	err = w.WriteAll(rows)
	if err != nil {
		return err
	}

	return nil
}

// create csv file from operating system
func createFile(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

// generate the necessary headers for csv file
func generateHeaders() []string {
	return []string{
		"Ementa",
		"Classificação",
	}
}

// returns a []string that compose the row in the csv file
func generateRow(result Summary.InferredDecision) []string {
	return []string{
		result.Summary,
		result.Class,
	}
}
