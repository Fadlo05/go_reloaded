package functions

func AToAn(s []string) []string {
	var res []string
	for i := 0; i < len(s); i++ {
		if s[i] == "a" || s[i] == "\na" || s[i] == "A" || s[i] == "\nA" {
			if i+1 < len(s) && CheckVowel(s[i+1]) {
				if s[i] == "a" {
					res = append(res, "an")
				} else if s[i] == "\na" {
					res = append(res, "\nan")
				} else if s[i] == "A" {
					res = append(res, "An")
				} else {
					res = append(res, "\nAn")
				}
				continue
			}
		}
		res = append(res, s[i])
	}
	return res
}
