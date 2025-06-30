package functions

import (
	"strconv"
	"strings"
)

func isFlag(flag string) bool {
	flag = strings.TrimSpace(flag)
	if !strings.HasPrefix(flag, "(") || !strings.HasSuffix(flag, ")") {
		return false
	}
	flag = flag[1 : len(flag)-1]
	parts := strings.Split(flag, ",")
	if len(parts) != 2 {
		return false
	}
	parts[0] = strings.TrimSpace(parts[0])
	parts[1] = strings.TrimSpace(parts[1])
	if parts[0] != "low" && parts[0] != "cap" && parts[0] != "up" {
		return false
	}
	num, err := strconv.Atoi(parts[1])
	if err != nil || num <= 0 {
		return false
	}
	return true
}
