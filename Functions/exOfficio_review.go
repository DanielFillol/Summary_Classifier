package Functions

func ExOfficioReview(text string, char int) bool {

	less17 := []string{"reexame necessário", "xame necessário"}
	more17 := []string{"reexame necessário"}

	return matchWords(text, char, less17, more17)

}
