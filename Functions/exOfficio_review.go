package Functions

func ExOfficioReview(text string, char int) bool {

	var less17 = []string{"reexame necessário", "xame necessário"}
	var more17 = []string{"reexame necessário"}

	ret := matchWords(text, char, less17, more17)

	return ret
}
