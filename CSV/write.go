package CSV

import (
	"encoding/csv"
	"github.com/Darklabel91/Summary_Classifier/Struct"
	"os"
	"path/filepath"
)

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func ExportCSV(nameFile string, nameFolder string, result []Struct.Infered_decision) error {
	var inferredDecision [][]string

	for i := 0; i < len(result); i++ {
		final := []string{result[i].Summary, result[i].Text, result[i].Class, result[i].Identifier, result[i].Court}
		inferredDecision = append(inferredDecision, final)
	}

	csvFile, err := create(nameFolder + "/" + nameFile + ".csv")
	if err != nil {
		return err
	}

	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)

	for _, infDec := range inferredDecision {
		_ = csvWriter.Write(infDec)
	}
	csvWriter.Flush()

	return nil
}
