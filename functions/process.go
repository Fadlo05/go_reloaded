package functions

import "fmt"

func Process(s []string) []string {
	result := []string{}

	for i := 0; i < len(s); i++ {
		if i == 0 && (s[i] == "(hex)" || s[i] == "(bin)" || s[i] == "(up)" || s[i] == "(low)" || s[i] == "(cap)") {
			continue
		}

		if i > 0 && (s[i] == "(hex)" || s[i] == "(hex)\n") {
			if HexToDecimal(result[len(result)-1]) == "error" {
				fmt.Println("this is not valid hex")
				return nil
			}
			result[len(result)-1] = HexToDecimal(result[len(result)-1])

		} else if i > 0 && (s[i] == "(bin)" || s[i] == "(bin)\n") {
			if BinToDecimal(result[len(result)-1]) == "error" {
				fmt.Println("this is not valid bin")
				return nil
			}
			result[len(result)-1] = BinToDecimal(result[len(result)-1])

		} else if i > 0 && (s[i] == "(up)" || s[i] == "(up)\n") {
			result[len(result)-1] = ToUpper(result[len(result)-1])
		} else if i > 0 && (s[i] == "(low)" || s[i] == "(low)\n") {
			result[len(result)-1] = ToLower(result[len(result)-1])
		} else if i > 0 && (s[i] == "(cap)" || s[i] == "(cap)\n") {
			result[len(result)-1] = Capitalize(result[len(result)-1])
		} else {
			result = append(result, s[i])
		}
	}

	return result
}
