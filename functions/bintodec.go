package functions

import (
	"strconv"
)

func BinToDecimal(s string) string {
	n, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return "error"
	}
	return strconv.Itoa(int(n))
}
