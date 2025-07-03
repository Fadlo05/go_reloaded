package functions

func AToAn(s []string) []string {
	var res []string
	for i := 0; i < len(s); i++ {
		if s[i] == "a" ||s[i] == "A" {
			if i+1 < len(s) && CheckVowel(s[i+1]) {
				if s[i] == "a" {
					res = append(res, "an")
				} else {
					res = append(res, "An")

				}
				continue
			}
		}
		res = append(res, s[i])
	}
	return res
}
