# Summary Classifier
Go Package build with the objective of classifying the legal court decisions writen in brazilian portuguese.

## Instal
``` go get -u github.com/Darklabel91/Summary_Classifier ```

## Data Struct
Return data can be ```bool```, ```string``` ou ```Infered_decision```, the last is constructed as: 
``` 
type InferredDecision struct {
	Summary string `json:"summary,omitempty"`
	Class   string `json:"classificcao,omitempty"`
}
```

- Summary: Summary of the legal decision you want to classify
- Class: The classification of the Summary provided. Can be one of the following:

## Categorias
- Prejudicado (Affected)
- Convertido em diligência (Diligence)
- Reexame necessário (ExOfficio)
- Parcial Provimento (Partial)
- Improvimento (Groundless)
- Provimento (Have Grounds)
- Casos não mapeados (Not mapped)

## Example

``` 
package main

import (
	"fmt"
	"github.com/Darklabel91/Summary_Classifier/SummaryCSV"
	"github.com/Darklabel91/Summary_Classifier/Summary"
)

func main() {

	//SINGLE USE OF THE CLASSIFIER
	summary := "PROCESSUAL CIVIL. AGRAVO INERNO NO RECURSO ESPECIAL. AÇÃO DE COMPENSAÇÃO POR DANO MORAL E REPARAÇÃO POR DANO MATERIAL. DANO MORAL. OCORRÊNCIA. LONGO ATRASO NA ENTREGA DE IMÓVEL. 1. Ação de compensação por dano moral e reparação por dano material. 2. Cabimento de compensação por danos morais em virtude de longo atraso na entrega de imóvel. Precedentes. 3. Agravo interno no recurso especial não provido."

	classification, err := Summary.Classify(summary)
	fmt.Println(classification.Summary, classification.Class, err)

	//READING A CSV WITH SUMMARY'S
	rawPath := "Summary_Classifier/CSV/csvTest.csv"
	separator := ';'
	resultFolder := "Result"

	err = SummaryCSV.SummaryClassifierCSV(rawPath, separator, resultFolder)
	if err != nil {
		fmt.Println(err)
	}

}
 ```
Return
``` 
PROCESSUAL CIVIL. AGRAVO INERNO NO RECURSO ESPECIAL. AÇÃO DE COMPENSAÇÃO POR DANO MORAL E REPARAÇÃO POR DANO MATERIAL. DANO MORAL. OCORRÊNCIA. LONGO ATRASO NA ENTREGA DE IMÓVEL. 1. Ação de compensação por dano moral e reparação por dano material. 2. Cabimento de compensação por danos morais em virtude de longo atraso na entrega de imóvel. Precedentes. 3. Agravo interno no recurso especial não provido. 

Improvimento 

<nil>

Files created
 ```

## Functions
- Classify(summary string) -> returns *[Infered_decision](https://pkg.go.dev/github.com/Darklabel91/Summary_Classifier/Summary#InferredDecision)* 
- ReturnSummaryClass(summary string, char int) -> returns one of seven possible classifications.
- SummaryClassifierCSV(rawFilePath string, separator rune, nameResultFolder string) -> return a CSV file in folder provide as parameter in *resultFolder*. The .csv passed in *rawFilePath* must contain a single column with summary text's.

## Disclaimer
The precision was calculated only with TJSP data. The precision was 96%;
