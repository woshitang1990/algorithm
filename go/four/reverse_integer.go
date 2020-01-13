package four

func reverse(x int) int {
	res := 0
	const INT_MAX = int(1<<31) - 1
	const INT_MIN = int(^INT_MAX)
	for {
		if x == 0 {
			break
		}

		pop := x % 10
		x = x / 10
		if res > INT_MAX/10 || (res == INT_MAX/10 && pop > 7) {
			res = 0
			break
		}
		if res < INT_MIN/10 || (res == INT_MIN/10 && pop < -8) {
			res = 0
			break
		}
		res = res*10 + pop
	}

	return res
}
