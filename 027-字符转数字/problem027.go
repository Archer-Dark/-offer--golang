package main

func myAtoi(s string) int {
	// 如果输入的是空字符串，直接返回0
	if s == "" {
		return 0
	}
	i, n := 0, len(s)
	// 忽略前导空格
	for i < n-1 {
		if s[i] == ' ' {
			i++
		} else {
			break
		}
	}
	// 检查非空格的第一位，并获得数字开头的s
	s0 := s[i:]
	if s[i] == '-' || s[i] == '+' {
		s = s[i+1:]
	} else if s[i] >= '0' && s[i] <= '9' {
		s = s[i:]
	} else {
		return 0
	}

	// 遍历s，并转化数字
	// 停止条件为：ch不是数字 或者 num超过最大长度
	num := 0
	for _, ch := range []byte(s) {
		ch -= '0'
		if ch > 9 || ch < 0 {
			break
		}

		num = num*10 + int(ch)
		if num >= 1<<31 {
			num = 1 << 31
			break
		}
	}

	// 有负号的安负号
	if s0[0] == '-' {
		num = -num
	} else {
		// 另一个边界
		if num >= 1<<31-1 {
			num = 1<<31 - 1
		}
	}

	return num
}
