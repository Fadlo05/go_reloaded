package functions

import "strings"

func ChangesWithN(s []string) []string {
	res := []string{}
	for i := 0; i < len(s); i++ {
		if strings.HasPrefix(s[i], "(cap") {
			num := getNumber(s[i])
			start := len(res) - num
			if start < 0 {
				start = 0
			}
			for j := start; j < len(res); j++ {
				res[j] = Capitalize(res[j])
			}
		} else if strings.HasPrefix(s[i], "(low") {
			num := getNumber(s[i])
			start := len(res) - num
			if start < 0 {
				start = 0
			}
			for j := start; j < len(res); j++ {
				res[j] = ToLower(res[j])
			}
		} else if strings.HasPrefix(s[i], "(up") {
			num := getNumber(s[i])
			start := len(res) - num
			if start < 0 {
				start = 0
			}
			for j := start; j < len(res); j++ {
				res[j] = ToUpper(res[j])
			}
		} else {
			res = append(res, s[i])
		}
	}
	return res
}
