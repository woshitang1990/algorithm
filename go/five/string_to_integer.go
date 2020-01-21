package five

import (
	"math"
)

func myAtoi(str string) int {
	if str == "" {
		return 0
	}
	sign := 1
	i := 0
SignLoop:
	for _, v := range str {
		switch v {
		case ' ':
			i++
		case '-':
			i++
			sign = -1
			break SignLoop
		case '+':
			i++
			break SignLoop
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			break SignLoop
		default:
			return 0
		}
	}

	str = str[i:]
	res := 0
ConvertLoop:
	for _, s := range str {
		if s < '0' || s > '9' {
			break ConvertLoop
		}

		res = res*10 + int(s-'0')

		if sign == 1 && res >= math.MaxInt32 {
			return math.MaxInt32
		}

		if sign == -1 && sign*res <= math.MinInt32 {
			return math.MinInt32
		}
	}

	return sign * res
}
