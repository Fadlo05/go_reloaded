package functions

func getNumber(s string) int{
	num := 0
	for _, c:= range s{
		if c == '-'{
			return 0
		}else if '0' <= c && c <= '9'{
			num = num*10 + int(c - '0')
		}
	}
	return num
}