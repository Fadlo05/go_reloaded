package functions

func IsPunctuation(s string) bool {
	return s == "." || s == "," || s == "!" || s == "?" || s == ":" || s == ";"
}

func Ponctuations(s []string) []string {
	var res []string
	i := 0
	for i < len(s) {
		if i+1 < len(s) && CheckPonc(s[i+1]) {
			res = append(res, s[i]+s[i+1])
			i += 2
		} else {
			res = append(res, s[i])
			i++
		}
	}
	return res
}
