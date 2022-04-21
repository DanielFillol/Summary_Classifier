package Functions

func Diligence(text string, char int) bool {

	less17 := []string{"conversão do julgamento em diligência", "convertido em diligência", "diligência"}
	more17 := []string{"acórdão em diligência", "conversão do julgamento em diligência", "conversão do feito em diligência", "convertido em diligência", "decisão em diligência", "julgalmento em diligência", "sentença em diligência"}

	return matchWords(text, char, less17, more17)

}
