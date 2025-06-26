package functions

func AToAn(s []string) []string {
	var res []string
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && s[i] == "a" && CheckVowel(s[i+1]) {
			res = append(res, "an")
		} else if i+1 < len(s) && s[i] == "A" && CheckVowel(s[i+1]) {
			res = append(res, "An")
		} else {
			res = append(res, s[i])
		}
	}
	return res
}
