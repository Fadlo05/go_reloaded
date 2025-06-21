package functions

import "strings"

func CheckPonc(s string) bool {
	if strings.HasPrefix(string(s[0]), ".") || strings.HasPrefix(string(s[0]), ",") || strings.HasPrefix(string(s[0]), "!") || strings.HasPrefix(string(s[0]), "?") || strings.HasPrefix(string(s[0]), ":") || strings.HasPrefix(string(s[0]), ";") {
		return true
	} else {
		return false
	}
}
