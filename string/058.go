package string

import "strings"

/*
给你一个字符串 s，由若干单词组成，单词之间用空格隔开。返回字符串中最后一个单词的长度。如果不存在最后一个单词，请返回 0 。
单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。


示例 1：
输入：s = "Hello World"
输出：5

示例 2：
输入：s = " "
输出：0
*/

// 方法一:
func LengthOfLastWord(s string) int {
	str := []byte(s)
	length := len(str)
	// 计数
	count := 0
	// 从后往前遍历
	for i := length - 1; i >= 0; i-- {
		// 每遍历一个计数就加一,当遇到空格返回
		if str[i] == ' ' {
			// 如果遇到空格的时候计数是0,表名字符以空格结尾,直接跳过
			if count == 0 {
				continue
			}
			break
		}
		count++
	}
	return count
}

// 方法二:
func LengthOfLastWord2(s string) int {
	str := strings.Trim(s," ")
	// 返回[]string
	str2 := strings.Split(str, " ")
	// Fields直接返回根据空格分隔字符串的[]string
	// str := strings.Fields(S)
	if len(str2) == 0 {
		return 0
	}
	return len(str2[len(str2)-1])
}
