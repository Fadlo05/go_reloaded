package functions

func AToAn(s []string) []string {
	var res []string
	i := 0
	for i < len(s) {
		if i+1 < len(s) && CheckVowel(s[i+1]) && s[i] == "a" {
			res = append(res, "an")
			i ++
		} else {
			res = append(res, s[i])
			i++
		}
	}
	return res
}
