# Summary Classifier
Projeto pensado para classificar as ementas de decisões extraídas das páginas dos tribunais de justiça de todo país.

## Instal
``` go get -u github.com/Darklabel91/Summary_Classifier ```

## Data Struct
Os dados de retorno podem ser ```bool```, ```string``` ou ```Infered_decision```, essa última é composta por:

``` 
type InferredDecision struct {
	Summary string `json:"summary,omitempty"`
	Class   string `json:"classificcao,omitempty"`
}
```

- Summary: Ementa ou trecho de decisão para classificar
- Class: A classificação feita (ver a seção "Categorias" abaixo)

## Categorias
- Prejudicado
- Convertido em diligência
- Reexame necessário
- Parcial Provimento
- Improvimento
- Provimento
- Casos não mapeados

## Example

``` 
package main

import (
	"fmt"
	"github.com/Darklabel91/Summary_Classifier/CSV"
	"github.com/Darklabel91/Summary_Classifier/Summary"
)

func main() {

	//SINGLE USE OF THE CLASSIFIER
	summary := "PROCESSUAL CIVIL. AGRAVO INERNO NO RECURSO ESPECIAL. AÇÃO DE COMPENSAÇÃO POR DANO MORAL E REPARAÇÃO POR DANO MATERIAL. DANO MORAL. OCORRÊNCIA. LONGO ATRASO NA ENTREGA DE IMÓVEL. 1. Ação de compensação por dano moral e reparação por dano material. 2. Cabimento de compensação por danos morais em virtude de longo atraso na entrega de imóvel. Precedentes. 3. Agravo interno no recurso especial não provido."

	classification, err := Summary.Classify(summary)
	fmt.Println(classification.Summary, classification.Class, err)

	//READING A CSV WITH SUMMARY'S
	rawPath := "/Users/danielfillol/Documents/GitHub/Summary_Classifier/CSV/csvTest.csv"
	separator := ';'
	resultFolder := "Result"

	err = CSV.SummaryClassifierCSV(rawPath, separator, resultFolder)
	if err != nil {
		fmt.Println(err)
	}

}
 ```
Retorno
``` 
PROCESSUAL CIVIL. AGRAVO INERNO NO RECURSO ESPECIAL. AÇÃO DE COMPENSAÇÃO POR DANO MORAL E REPARAÇÃO POR DANO MATERIAL. DANO MORAL. OCORRÊNCIA. LONGO ATRASO NA ENTREGA DE IMÓVEL. 1. Ação de compensação por dano moral e reparação por dano material. 2. Cabimento de compensação por danos morais em virtude de longo atraso na entrega de imóvel. Precedentes. 3. Agravo interno no recurso especial não provido. 

Improvimento 

<nil>

Files created
 ```

## Functions
- Classify(summary string) -> retorna *[Infered_decision](https://pkg.go.dev/github.com/Darklabel91/Summary_Classifier/Summary#InferredDecision)* 
- ReturnSummaryClass(summary string, char int) -> retorna uma das sete categorias de classificação possíveis.
- SummaryClassifierCSV(rawFilePath string, separator rune, nameResultFolder string) -> retorna um CSV para uma pasta do projeto com o nome apontado em *resultFolder*. Para utilizar a função basta apontar o caminho do CSV (deve possuir uma única coluna com as ementas)

## Disclaimer
Esse classificador foi testado, até o momento, apenas com ementas dos julgados do segundo grau do TJSP (Tribunal de Justiça de São Paulo) com uma assertividade de 96%, de qualquer modo, use com cautela.

## Considerações
A) Esse projeto foi criado de forma voluntária, você pode contribuir de qualquer modo. Se encontrar uma falha, não hesite em criar um “issue” ou  procure corrigir você mesma(o) o erro e dar um “pull request”.

B) use os dados baixados para agregar valor, como por exemplo, para realizar análises ou publicar artigos, fazer inferências, elaborar recomendações aos poderes públicos etc. Baixar esses dados para reproduzi-los em sua página web é tirar proveito do trabalho alheio, mesmo sendo esses dados públicos.
