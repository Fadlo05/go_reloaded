package functions

import (
	"strings"
)

func isPunct(c byte) bool {
	return c == '.' || c == ',' || c == '!' || c == '?' || c == ':' || c == ';'
}

func Tokenizer(text string) []string {
	parts := strings.Split(text, " ")

	str := []string{}

	for _, word := range parts {
		if len(word) == 0 {
			continue
		}

		if word[0] == '\'' && word[len(word)-1] == '\'' && len(word) > 2 {
			str = append(str, "'", word[1:len(word)-1], "'")
			continue
		}

		for len(word) > 0 && isPunct(word[0]) {
			str = append(str, string(word[0]))
			word = word[1:]
		}
		if len(word) == 0 {
			continue
		}

		end := len(word)
		for end > 0 && isPunct(word[end-1]) {
			end--
		}

		str = append(str, word[:end])

		if end < len(word) {
			str = append(str, word[end:])
		}
	}

	return str
}
