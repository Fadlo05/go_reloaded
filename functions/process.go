package functions

import (
	"fmt"
	"strings"
)

func lastValidIndex(result []string) int {
	for i := len(result) - 1; i >= 0; i-- {
		if result[i] != "\n" {
			return i
		}
	}
	return -1
}

func Process(s []string) []string {
	result := []string{}

	for i := 0; i < len(s); i++ {
		// Ne rien faire de spécial ici pour "\n", on le garde

		if i == 0 && (s[i] == "(hex)" || s[i] == "(bin)" || s[i] == "(up)" || s[i] == "(low)" || s[i] == "(cap)") {
			continue
		}

		if i > 0 && (s[i] == "(hex)" || s[i] == "(hex)\n" || s[i] == "(hex) ") {
			// Ignorer les \n ici aussi (mais résultat l’a déjà stocké)
			idx := lastValidIndex(result)
			if idx == -1 || HexToDecimal(result[idx]) == "error" {
				fmt.Println("this is not valid hex")
				return nil
			}
			result[idx] = HexToDecimal(result[idx])

		} else if i > 0 && (s[i] == "(bin)" || s[i] == "(bin)\n" || s[i] == "(bin) ") {
			idx := lastValidIndex(result)
			if idx == -1 || BinToDecimal(result[idx]) == "error" {
				fmt.Println("this is not valid bin")
				return nil
			}
			result[idx] = BinToDecimal(result[idx])

		} else if i > 0 && (s[i] == "(up)" || s[i] == "(up)\n" || s[i] == "(up) ") {
			idx := lastValidIndex(result)
			if idx != -1 {
				result[idx] = ToUpper(result[idx])
			}

		} else if i > 0 && (s[i] == "(low)" || s[i] == "(low)\n" || s[i] == "(low) ") {
			idx := lastValidIndex(result)
			if idx != -1 {
				result[idx] = ToLower(result[idx])
			}

		} else if i > 0 && (s[i] == "(cap)" || s[i] == "(cap)\n" || s[i] == "(cap) ") {
			idx := lastValidIndex(result)
			if idx != -1 {
				result[idx] = Capitalize(result[idx])
			}

		} else if strings.HasPrefix(s[i], "(cap") && strings.HasSuffix(s[i], ")") && isFlag(s[i]) {
			num := getNumber(s[i])
			count := 0
			for j := len(result) - 1; j >= 0 && count < num; j-- {
				if result[j] != "\n" {
					result[j] = Capitalize(result[j])
					count++
				}
			}

		} else if strings.HasPrefix(s[i], "(low") && strings.HasSuffix(s[i], ")") && isFlag(s[i]) {
			num := getNumber(s[i])
			count := 0
			for j := len(result) - 1; j >= 0 && count < num; j-- {
				if result[j] != "\n" {
					result[j] = ToLower(result[j])
					count++
				}
			}

		} else if strings.HasPrefix(s[i], "(up") && strings.HasSuffix(s[i], ")") && isFlag(s[i]) {
			num := getNumber(s[i])
			count := 0
			for j := len(result) - 1; j >= 0 && count < num; j-- {
				if result[j] != "\n" {
					result[j] = ToUpper(result[j])
					count++
				}
			}

		} else {
			result = append(result, s[i]) // on garde aussi les "\n" ici
		}
	}

	return result
}
