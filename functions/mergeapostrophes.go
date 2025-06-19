package functions

import "strings"

func MergeApostrophes(s []string) []string {
	res := []string{}
	str := ""
	inQuote := false
	for i := 0; i < len(s); i++ {
		if s[i] == "'" {
			if !inQuote {
				inQuote = true
				str = "'"
			} else {
				inQuote = false
				inner := strings.TrimSpace(str[1:])
				str = "'" + inner 
				str += "'"
				res = append(res, str)
				str = ""
			}
		} else if inQuote {
			if str == "'" {
				str += s[i]
			} else {
				str += " " + s[i]
			}
		} else {
			res = append(res, s[i])
		}
	}
	if inQuote && str != "" {
		inner := strings.TrimSpace(str[1:])
        str = "'" + inner
        res = append(res, str)
	}
	return res
}
