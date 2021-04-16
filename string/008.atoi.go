package string

import "strings"

/*
示例 4：
输入：s = "words and 987"
输出：0
解释：
第 1 步："words and 987"（当前没有读入字符，因为没有前导空格）
         ^
第 2 步："words and 987"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
         ^
第 3 步："words and 987"（由于当前字符 'w' 不是一个数字，所以读入停止）
         ^
解析得到整数 0 ，因为没有读入任何数字。
由于 0 在范围 [-231, 231 - 1] 内，最终结果为 0 。
示例 5：

输入：s = "-91283472332"
输出：-2147483648
解释：
第 1 步："-91283472332"（当前没有读入字符，因为没有前导空格）
         ^
第 2 步："-91283472332"（读入 '-' 字符，所以结果应该是负数）
          ^
第 3 步："-91283472332"（读入 "91283472332"）
                     ^
解析得到整数 -91283472332 。
由于 -91283472332 小于范围 [-231, 231 - 1] 的下界，最终结果被截断为 -231 = -2147483648 。
*/

func MyAtoi(s string) int {
	num, max, min := 0, 1<<31-1, -1<<31
	sign,abs := clean(s)
	for _, v := range abs {
		num = num*10 + int(v-'0')
		if sign == 1 && num > max {
			return max
		}
		if sign == -1 && sign*num < min {
			return min
		}
	}
	return sign*num
}
// 修剪字符串,返回符号和纯数字
func clean(s string) (sign int, abs string) {
	s = strings.TrimSpace(s) // 首先去括号
	if s == "" { 
		return
	}
	switch s[0] {
	case '+':
		sign, abs = 1, s[1:]
	case '-':
		sign, abs = -1, s[1:]
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		sign, abs = 1, s
	default:
		abs = ""
		return
	}
	for i, v := range abs { // 剔除后面的字母
		if v < '0' || v > '9' {
			abs = abs[:i]
			break // 必须break,
		}
	}
	return
}
