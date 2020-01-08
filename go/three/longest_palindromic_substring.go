package three

/**
 *
 * 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
 */

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	// 第一步：预处理字符串
	newStr := "^"
	for _, v := range s {
		newStr = newStr + "#" + string(v)
	}
	newStr += "#$"

	// 第二步：计算最长回文数组p，起始索引，最长回文半径
	// 预处理后的字符串长度
	l := len(newStr)
	// 最长回文数组
	p := make([]int, l, l)
	// 当前中心点位置
	id := 0
	// 当前中心点最右端位置
	mx := 0
	// 最长回文子串长度
	maxLength := -1
	// 最长回文子串中心位置索引
	index := 0
	for i := 1; i < l-1; i++ {
		if mx > i {
			j := 2*id - i
			if p[j] > mx-i {
				p[i] = mx - i
			} else {
				p[i] = p[j]
			}
		} else {
			p[i] = 1
		}

		// 继续延伸，判断是否超出当前回文子串范围
		for {
			if newStr[i+p[i]] == newStr[i-p[i]] {
				p[i]++
				continue
			}

			break
		}

		// 如果回文子串右边界超过了mx，则需要更新mx和id的值
		if mx < p[i]+i {
			mx = p[i] + i
			id = i
		}

		// 如果回文子串的长度大于maxLength, 则更新maxLength和index的值
		if maxLength < p[i]-1 {
			maxLength = p[i] - 1
			index = i
		}
	}

	// 第三步：截取字符串，输出结果
	start := (index - maxLength) / 2
	return s[start : start+maxLength]
}
