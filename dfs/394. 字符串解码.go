package dfs

import (
	"container/list"
	"strconv"
	"strings"
)

/*
给定一个经过编码的字符串，返回它解码后的字符串。
编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。


示例 1：
输入：s = "3[a]2[bc]"
输出："aaabcbc"

示例 2：
输入：s = "3[a2[c]]"
输出："accaccacc"

示例 3：
输入：s = "2[abc]3[cd]ef"
输出："abcabccdcdcdef"

示例 4：
输入：s = "abc3[cd]xyz"
输出："abccdcdcdxyz"
*/

func DecodeString(s string) string {
	nstark := []int{}
	stark := []string{}
	num := 0
	str := ""
	// 遍历字符串
	for i := 0; i < len(s); i++ {
		switch {
		// 当遇到数字算出需要倍数
		case s[i] >= '0' && s[i] <= '9':
			n, _ := strconv.Atoi(string(s[i]))
			num = num*10 + n
			// 遇到'['倍数和字符入栈
		case s[i] == '[':
			nstark = append(nstark, num)
			num = 0
			stark = append(stark, str)
			str = ""
			// 遇到']'栈顶的倍数和字符出栈拼接字符串
		case s[i] == ']':
			ntop := nstark[len(nstark)-1]
			nstark = nstark[:len(nstark)-1]
			stop := stark[len(stark)-1]
			stark = stark[:len(stark)-1]
			str = stop + strings.Repeat(str, ntop)
			// 遇到字符
		default:
			str += string(s[i])
		}
	}
	return str
}

// list库改写
func DecodeString2(s string) string {
	type pair struct {
		bs  int
		str string
	}
	stark := list.New()
	num, res := 0, ""
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] >= '0' && s[i] <= '9':
			num = num*10 + int(s[i]-'0')
		case s[i] == '[':
			stark.PushBack(pair{num, res})
			num, res = 0, ""
		case s[i] == ']':
			top := stark.Back().Value.(pair)
			stark.Remove(stark.Back())
			res = top.str + strings.Repeat(res, top.bs)
		default:
			res += string(s[i])
		}
	}
	return res
}
