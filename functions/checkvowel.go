package functions

import "strings"

func CheckVowel(s string) bool {
	if strings.HasPrefix(string(s[0]), "a") || strings.HasPrefix(string(s[0]), "e") || strings.HasPrefix(string(s[0]), "i") || strings.HasPrefix(string(s[0]), "o") || strings.HasPrefix(string(s[0]), "u") || strings.HasPrefix(string(s[0]), "y") || strings.HasPrefix(string(s[0]), "h") {
		return true
	} else {
		return false
	}
}
