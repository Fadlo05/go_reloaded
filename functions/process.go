package functions

import "fmt"

func Process(s []string) []string {
	result := []string{}

	for i := 0; i < len(s); i++ {
		if i == 0 && (s[i] == "(hex)" || s[i] == "(bin)" || s[i] == "(up)" || s[i] == "(low)" || s[i] == "(cap)") {
			continue
		}

		if s[i] == "(hex)" && i > 0 {
			if HexToDecimal(result[len(result)-1]) == "error" {
				fmt.Println("this is not valid hex")
				return nil
			}
			result[len(result)-1] = HexToDecimal(result[len(result)-1])

		} else if s[i] == "(bin)" && i > 0 {
			if BinToDecimal(result[len(result)-1]) == "error" {
				fmt.Println("this is not valid bin")
				return nil
			}
			result[len(result)-1] = BinToDecimal(result[len(result)-1])

		} else if s[i] == "(up)" && i > 0 {
			result[len(result)-1] = ToUpper(result[len(result)-1])
		} else if s[i] == "(low)" && i > 0 {
			result[len(result)-1] = ToLower(result[len(result)-1])
		} else if s[i] == "(cap)" && i > 0 {
			result[len(result)-1] = Capitalize(result[len(result)-1])
		} else {
			result = append(result, s[i])
		}
	}

	return result
}
