package functions

import (
	"fmt"
	"strings"
)

func Process(s []string) []string {
	result := []string{}

	for i := 0; i < len(s); i++ {
		if i == 0 && (s[i] == "(hex)" || s[i] == "(bin)" || s[i] == "(up)" || s[i] == "(low)" || s[i] == "(cap)") {
			continue
		}

		if i > 0 && (s[i] == "(hex)" || s[i] == "(hex)\n" || s[i] == "(hex) ") {
			if HexToDecimal(result[len(result)-1]) == "error" {
				fmt.Println("this is not valid hex")
				return nil
			}
			result[len(result)-1] = HexToDecimal(result[len(result)-1])

		} else if i > 0 && (s[i] == "(bin)" || s[i] == "(bin)\n" || s[i] == "(bin) ") {
			if BinToDecimal(result[len(result)-1]) == "error" {
				fmt.Println("this is not valid bin")
				return nil
			}
			result[len(result)-1] = BinToDecimal(result[len(result)-1])

		} else if i > 0 && (s[i] == "(up)" || s[i] == "(up)\n" || s[i] == "(up) ") {
			result[len(result)-1] = ToUpper(result[len(result)-1])
		} else if i > 0 && (s[i] == "(low)" || s[i] == "(low)\n" || s[i] == "(low) ") {
			result[len(result)-1] = ToLower(result[len(result)-1])
		} else if i > 0 && (s[i] == "(cap)" || s[i] == "(cap)\n" || s[i] == "(cap) ") {
			result[len(result)-1] = Capitalize(result[len(result)-1])
		} else if strings.HasPrefix(s[i], "(cap") && strings.HasSuffix(s[i], ")") && isFlag(s[i]) {
			num := getNumber(s[i])
			start := len(result) - num
			if start < 0 {
				start = 0
			}
			for j := start; j < len(result); j++ {
				result[j] = Capitalize(result[j])
			}
		} else if strings.HasPrefix(s[i], "(low") && strings.HasSuffix(s[i], ")") && isFlag(s[i]) {
			num := getNumber(s[i])
			start := len(result) - num
			if start < 0 {
				start = 0
			}
			for j := start; j < len(result); j++ {
				result[j] = ToLower(result[j])
			}
		} else if strings.HasPrefix(s[i], "(up") && strings.HasSuffix(s[i], ")") && isFlag(s[i]) {
			num := getNumber(s[i])
			start := len(result) - num
			if start < 0 {
				start = 0
			}
			for j := start; j < len(result); j++ {
				result[j] = ToUpper(result[j])
			}
		} else {
			result = append(result, s[i])
		}
	}

	return result
}
