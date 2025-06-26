package functions

func Clean(input string) []string {
	// Convertit la chaîne en []rune pour indexer par caractère logique (rune)
	runes := []rune(input)
	var words []string
	start := 0
	isInsideParentheses := false

	for i, r := range runes {
		// if r == '(' {
		// 	isInsideParentheses = true
		// } else if r == ')' {
		// 	isInsideParentheses = false
		// }

		// Détection de fin de mot : espace hors parenthèses
		if r == ' ' && !isInsideParentheses {
			if start < i {
				words = append(words, string(runes[start:i]))
			}
			start = i + 1
		}
	}

	// Ajoute le dernier mot s'il reste quelque chose
	if start < len(runes) {
		words = append(words, string(runes[start:]))
	}

	return words
}
