package functions

func Process(s []string) []string{
	result := [] string {}
	for i := 0; i < len(s); i++ {
		if s[i + 1] == "(hex)" && i > 0{
			s[i] = HexToDecimal(s[i -1])
		}
	}
}