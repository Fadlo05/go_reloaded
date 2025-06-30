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
					if i < len(runes) && runes[i] == ' ' {
						i++
					}
					start = i
					continue
				}
			}
		}
		if runes[i] == ' ' {
			if start < i {
				words = append(words, string(runes[start:i]))
			}
			start = i + 1
		}
		i++
	}
	if start < len(runes) {
		words = append(words, string(runes[start:]))
	}
	return words
}
