package functions

func Clean(input string) []string {
	runes := []rune(input)
	var words []string
	start := 0
	i := 0
	for i < len(runes) {
	if runes[i] == '(' {
		j := i
		for j < len(runes) && runes[j] != ')' {
			j++
		}
		if j < len(runes) && runes[j] == ')' {
			block := string(runes[i : j+1])
			if isFlag(block) {
				words = append(words, block)
				i = j + 1
				if i < len(runes) && (runes[i] == ' ' || runes[i] == '\n') {
					i++
				}
				start = i
				continue
			}
		}
	}
	if runes[i] == ' ' || runes[i] == '\n' {
		if start < i {
			word := string(runes[start:i])
			if runes[i] == '\n' {
				word += "\n"
			}
			words = append(words, word)
		}
		start = i + 1
	}
	i++
}
if start < len(runes) {
	word := string(runes[start:])
	words = append(words, word)
}

	return words
}
