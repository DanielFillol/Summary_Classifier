package CSV

import (
	"encoding/csv"
	"fmt"
	"github.com/Darklabel91/Summary_Classifier/Error"
	"github.com/Darklabel91/Summary_Classifier/Struct"
	"os"
	"path/filepath"
)

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	fmt.Println("Criou .csv")
	return os.Create(p)
}

func ExportCSV(nameFile string, nameFolder string, result []Struct.Infered_decision) {
	empData := [][]string{}

	for i := 0; i < len(result); i++ {
		final := []string{result[i].Summary, result[i].Text, result[i].Class, result[i].Identifier, result[i].Court}
		empData = append(empData, final)
	}

	csvFile, _ := create(nameFolder + "/" + nameFile + ".csv")
	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range empData {
		_ = csvwriter.Write(empRow)
	}
	csvwriter.Flush()
	err := csvFile.Close()
	Error.CheckError(err)
}
