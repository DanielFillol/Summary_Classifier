package Functions

func Affected(text string, char int) bool {

	var less17 = []string{"dicado", "extinção da punibilidade", "icado", "perda do objeto", "prejudicad", "prejudicado", "prejudicada", "perda superveniente do objeto", "prejudicando o exame"}
	var more17 = []string{"extinção da punibilidade", "perda do objeto", "prejudicad", "prejudicado", "prejudicada", "perda superveniente do objeto", "prejudicando o exame"}

	ret := matchWords(text, char, less17, more17)

	return ret
}
