package CSV

import (
	"encoding/csv"
	"github.com/Darklabel91/Summary_Classifier/Struct"
	"os"
)

func ReadCsvFile(filePath string, separator rune) ([]Struct.Raw_decision, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		return []Struct.Raw_decision{}, err
	}

	defer csvFile.Close()

	csvR := csv.NewReader(csvFile)
	csvR.Comma = separator

	csvData, err := csvR.ReadAll()
	if err != nil {
		return []Struct.Raw_decision{}, err
	}

	var data []Struct.Raw_decision

	for _, line := range csvData {
		emp := Struct.Raw_decision{
			Summary:    line[0],
			Identifier: line[1],
			Court:      line[2],
		}

		data = append(data, emp)
	}

	return data, nil
}
