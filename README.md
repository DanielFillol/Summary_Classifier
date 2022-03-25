# Summary Classifier
Projeto pensado para classificar as ementas de decisões extraídas das páginas dos tribunais de justiça de todo país.

## Instal
``` go get -u github.com/Darklabel91/Summary_Classifier ```

## Data Struct
Os dados de retorno podem ser ```bool```, ```string``` ou ```Infered_decision```, essa última é composta por:

``` 
type Infered_decision struct {
	Summary    string `json:"summary,omitempty"`
	Text       string `json:"text,omitempty"`
	Class      string `json:"classificcao,omitempty"`
	Identifier string `json:"Identifier,omitempty"`
	Court      string `json:"Court,omitempty"`
}
```

- Summary: Ementa ou trecho de decisão para classificar
- Text: Trecho de texto selecionado pelo classificador
- Class: A classificação feita (ver a seção "Categorias" abaixo)
- Identifier: Um identificador. É recomendado que seja utilizado o número do CNJ que a ementa está conectada para facilitar o cruzamento de dados
- Court: O tribunal do qual a ementa foi extraída

## Categorias
- Prejudicado
- Convertido em diligência
- Reexame necessário
- Parcial Provimento
- Improvimento
- Provimento
- Sem Ementa
- Casos não mapeados

## Example

``` 
package main

import (
	"fmt"
	"github.com/Darklabel91/Summary_Classifier"
)

func main() {
	
	//SINGLE USE OF THE CLASSIFIER

	summary := "PROCESSUAL CIVIL. AGRAVO INERNO NO RECURSO ESPECIAL. AÇÃO DE COMPENSAÇÃO POR DANO MORAL E REPARAÇÃO POR DANO MATERIAL. DANO MORAL. OCORRÊNCIA. LONGO ATRASO NA ENTREGA DE IMÓVEL. 1. Ação de compensação por dano moral e reparação por dano material. 2. Cabimento de compensação por danos morais em virtude de longo atraso na entrega de imóvel. Precedentes. 3. Agravo interno no recurso especial não provido."
	id := "0"
	court := "STJ"

	test := Summary_Classifier.SummaryClassifier(summary, id, court)
	fmt.Println(test.Class)

	//READING A CSV WITH SUMMARY'S

	rawPath := "/Users/danielfillol/Desktop/Decisioes.csv"
	separator := ';'
	resultFolder := "Result"

	Summary_Classifier.SummaryClassifierCSV(rawPath, separator, resultFolder)

}

 ```
Retorno
``` 
Improvimento
Files created
 ```

## Functions

Main Function:
- Decision_Classifier(summary string, identifier string, court string)  ->  retorna uma *Infered_decision* necessitantando da ementa, identificador, tribunal. Essa função faz um uso em laço da *ClassDecision*, iniciando com 16 caracteres, até a totalidade de caracteres da ementa para classificar o texto.
- SummaryClassifierCSV(rawPath string, separator rune, resultFolder string)-> retorna um CSV para uma pasta do projeto com o nome apontado em *resultFolder*. Para utilizar a função basta apontar o caminho do CSV (que deve ter a sequência de colunas {id, identifier, decision, court}) e o separador (';' ',' etc..)

Decision Function:
- ClassDecision(summary string, identifier string, court string, char int)  ->  retorna uma *Infered_decision* necessitantando da ementa, identificador, tribunal e número de caracteres a serem analisados na ementa (de trás para frente).

Suport Functions:
- Affected(text string, char int)        ->  retorna true para uma ementa *prejudicada*
- Diligence(text string, char int)       ->  retorna true para uma ementa *convertida em diligência*
- ExOfficioReview(text string, char int) ->  retorna true para uma ementa *reexame necessário*
- Partial(text string, char int)         ->  retorna true para uma ementa *parcial provimento*
- Groundless(text string, char int)      ->  retorna true para uma ementa *improvimento*
- HasGround(text string, char int)       ->  retorna true para uma ementa *provimento* 

## CSV return files
Serão retornados 9 arquivos csv apartados correspondetes a uma classificação mencionada acima, sendo esses:

- affected.csv
- diligence.csv
- exOfficioReview.csv
- partial.csv 
- groundless.csv 
- hasGrounds.csv 
- noSummary.csv
- notMapped.csv 
- totalInfered.csv <> **Compilado de todos os documentos acima**

Caso não exista nenhum elemento de uma dada categoria o arquivo não será retornado.

## Disclaimer
Esse é um work in progress do classificador, ainda não sabemos o % de efetividade desse classificador para as ementas existentes nos tribunais de justiça. Use por sua conta em risco.

## Considerações
A) Esse projeto foi criado de forma voluntária, você pode contribuir de qualquer modo. Se encontrar uma falha, não hesite em criar um “issue” ou  procure corrigir você mesma(o) o erro e dar um “pull request”.

B) use os dados baixados para agregar valor, como por exemplo, para realizar análises ou publicar artigos, fazer inferências, elaborar recomendações aos poderes públicos etc. Baixar esses dados para reproduzi-los em sua página web é tirar proveito do trabalho alheio, mesmo sendo esses dados públicos.
