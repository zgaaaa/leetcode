package string

import "strings"

/*
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:
输入: "A man, a plan, a canal: Panama"
输出: true

示例 2:
输入: "race a car"
输出: false
*/

// 方法一：
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	a := 0
	b := len(s) - 1
	for a < b {
		if !((s[a] >= '0' && s[a] <= '9') || (s[a] >= 'a' && s[a] <= 'z')) {
			a++
			continue
		}
		if !((s[b] >= '0' && s[b] <= '9') || (s[b] >= 'a' && s[b] <= 'z')) {
			b--
			continue
		}
		if s[a] != s[b] {
			return false
		}
		a++
		b--

	}
	return true	
}
//方法二:
func IsPalindrome2(s string) bool {
	s = strings.ToLower(s)
	str := []byte(s)
	a := 0
	for i := 0; i < len(str); i++ {
		if (str[i] >= '0' && str[i] <= '9') || (str[i] >= 'a' && str[i] <= 'z') {
			str[a] = str[i]
			a++
		}

	}
	for i := 0; i < a; i++ {
		if str[a-i-1] != str[i] {
			return false
		}
	}
	return true
}