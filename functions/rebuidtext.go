package functions

func RebuildText(s []string) string{
	res := ""
	for i := 0; i < len(s); i++ {
		res += s[i] + " "
	}
	return  res
}