package string

import (
	"strconv"
	"strings"
)

/*
字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。
比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变短，则返回原先的字符串。
你可以假设字符串中只包含大小写英文字母（a至z）。

示例1:
 输入："aabcccccaaa"
 输出："a2b1c5a3"

示例2:
 输入："abbccd"
 输出："abbccd"
 解释："abbccd"压缩后为"a1b2c2d1"，比原字符串长度更长。
*/

func CompressString(S string) string {
	slen := len(S)
	// res := make([]byte,0,slen)
	var res strings.Builder
	for i := 0; i < slen; i++ {
		count := 1 // 计数器
		for i+1 < slen && S[i] == S[i+1] {
			count++ // 当前字符和下一个字符相同时计数器加一
			i++     // 下标同时移动
		}
		// 字符串拼接
		// res = append(res, S[i])
		// res = append(res, []byte(strconv.Itoa(count))...)
		res.WriteByte(S[i])
		res.WriteString(strconv.Itoa(count))
		// 如果当前长度大于原字符串长的就不需要继续运算了,直接返回
		if slen <= len(res.String()) {
			return S
		}
	}
	return string(res.String())
}
