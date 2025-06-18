package functions

func ToUpper(s string) string {
	str := []rune(s)
	for i := 0; i < len(str); i++ {
		if 'a' <= str[i] && str[i] <= 'z' {
			str[i] = str[i] - 32
		}
	}
	return string(str)
}
