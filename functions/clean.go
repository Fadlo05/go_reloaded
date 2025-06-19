package functions

func Clean(input string) []string {
	startIndex := 0
	words := []string{}
	isInsideParentheses := false
	for i, char := range input {
		if char == '(' {
			isInsideParentheses = true
		} else if char == ')' {
			isInsideParentheses = false
		}
		if char != ' ' && i != len(input)-1 && !isInsideParentheses {
			startIndex++
		} else if isInsideParentheses {
			startIndex++
		}
		if char == ' ' && startIndex != 0 && !isInsideParentheses {
			wordStart := i - startIndex
			words = append(words, input[wordStart:i])
			startIndex = 0

		} else if i == len(input)-1 && !isInsideParentheses && char != ' ' {
			wordStart := i - startIndex
			words = append(words, input[wordStart:])
		}
	}
	return words
}
