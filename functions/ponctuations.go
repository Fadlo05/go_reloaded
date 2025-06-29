package functions

import (
	"strings"
	"unicode"
)

func noLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func Ponctuations(s []string) string {
	var res []string
	par := ""

	// Première étape : regrouper ponctuations entre elles
	for i := 0; i < len(s); i++ {
		if CheckPonc(s[i]) {
			par += s[i]
		} else {
			if par != "" {
				res = append(res, par)
				par = ""
			}
			res = append(res, s[i])
		}
	}
	if par != "" {
		res = append(res, par)
	}

	// Deuxième étape : concaténer ponctuations au mot précédent
	var r []string
	for _, tok := range res {
		if len(r) > 0 &&
			strings.ContainsAny(tok, ".,!?:;") &&
			noLetter(tok) {
			r[len(r)-1] += tok
		} else {
			r = append(r, tok)
		}
	}

	resStr := strings.Join(r , " ")

	return resStr
}
