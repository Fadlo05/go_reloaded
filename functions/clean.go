package functions

func Clean(input string) []string {
	runes := []rune(input)
	var words []string
	start := 0
	i := 0

	for i < len(runes) {
		// Si on détecte une parenthèse ouvrante
		if runes[i] == '(' {
			j := i
			// Chercher la parenthèse fermante correspondante
			for j < len(runes) && runes[j] != ')' {
				j++
			}

			// Si une parenthèse fermante a été trouvée
			if j < len(runes) && runes[j] == ')' {
				block := string(runes[i : j+1])

				// Vérifier si c'est un flag valide
				if isFlag(block) {
					words = append(words, block)
					i = j + 1

					// Sauter l'espace juste après le flag s’il y en a
					if i < len(runes) && runes[i] == ' ' {
						i++
					}
					start = i
					continue
				}
			}
		}

		// Détection des mots hors parenthèses
		if runes[i] == ' ' {
			if start < i {
				words = append(words, string(runes[start:i]))
			}
			start = i + 1
		}

		i++
	}

	// Ajouter le dernier mot
	if start < len(runes) {
		words = append(words, string(runes[start:]))
	}

	return words
}
