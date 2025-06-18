package functions

import (
	"strconv"
)

func HexToDecimal(s string) string {
	n, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return ""
	}
	return strconv.Itoa(int(n))
}
