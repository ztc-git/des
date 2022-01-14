package des

func InitReplacement(text string) string {
	replaceText := ""
	initialReplaceMatrix := GetInitialReplaceMatrix()
	for i := 0; i < len(initialReplaceMatrix); i++ {
		replaceText += string(text[initialReplaceMatrix[i] - 1])
	}

	return replaceText
}

func ReverseReplacement(text string) string {
	replaceText := ""
	reverseReplaceMatrix := GetReverseReplace()
	for i := 0; i < len(reverseReplaceMatrix); i++ {
		replaceText += string(text[reverseReplaceMatrix[i] - 1])
	}

	return replaceText
}