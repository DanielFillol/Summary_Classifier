package Functions

func Affected(text string, char int) bool {

	less17 := []string{"dicado", "extinção da punibilidade", "icado", "perda do objeto", "prejudicad", "prejudicado", "prejudicada", "perda superveniente do objeto", "prejudicando o exame"}
	more17 := []string{"extinção da punibilidade", "perda do objeto", "prejudicad", "prejudicado", "prejudicada", "perda superveniente do objeto", "prejudicando o exame"}

	return matchWords(text, char, less17, more17)

}
