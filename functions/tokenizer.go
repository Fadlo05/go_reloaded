package functions

func isPunct(c byte) bool {
	return c == '.' || c == ',' || c == '!' || c == '?' || c == ':' || c == ';'
}

func isQuote(c byte) bool {
	return c == '\''
}

func Tokenizer(words []string) []string {
	str := []string{}

	for _, word := range words {
		if len(word) == 0 {
			continue
		}

		if word[0] == '\'' && word[len(word)-1] == '\'' && len(word) > 2 {
			str = append(str, "'")
			str = append(str, word[1:len(word)-1])
			str = append(str, "'")
			continue
		}
		start := 0
		for start < len(word) && isQuote(word[start]){
			start++
		}
		if start > 0 {
			str = append(str, word[:start])
			word = word[start:]
		}
		start = 0
		for start < len(word) && isPunct(word[start]){
			start++
		}
		if start > 0 {
			str = append(str, word[:start])
			word = word[start:]
		}
		end := len(word)
		for end > 0 && (isPunct(word[end-1]) || isQuote(word[end-1])) {
			end--
		}
		if end > 0 {
			str = append(str, word[:end])
		}
		if end < len(word) {
			str = append(str, word[end:])
		}
	}
	return str
}
