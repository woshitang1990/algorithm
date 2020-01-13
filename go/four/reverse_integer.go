package four

/**
 *
 * 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
 * 注意:
 * 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。
 * 请根据这个假设，如果反转后整数溢出那么就返回 0
 * 链接：https://leetcode-cn.com/problems/reverse-integer/
 */

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
