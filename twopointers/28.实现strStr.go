package twopointers

/*
实现 strStr() 函数。
给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。
如果不存在，则返回  -1。

示例 1:
输入: haystack = "hello", needle = "ll"
输出: 2

示例 2:
输入: haystack = "aaaaa", needle = "bba"
输出: -1
*/

// 双指针法:
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	i, j := 0, 0
	index := -1 // 用于保存目标串第一次出现的位置
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			if index == -1 { // 表名目标串的第一个字符第一次出现
				index = i
			}
			if j == len(needle)-1 { // 如果j指向needle最后一位,就说明找到了
				return index
			}
			j++
		} else { // 如果不相等
			if index != -1 { // 而且当前查找的不是目标串的第一个字符
				// 指针回退
				i = index  // i 回退到index
				index = -1 // index回退到-1
				j = 0      // j回退到0,指向needle的第一位
			}
		}
		i++
	}
	return -1
}

// 双指针
func strStr2(haystack, needle string) int {
	if haystack == "" {
		return 0
	}
	l1, l2 := len(haystack), len(needle)
	for i := 0; i <= l1-l2; i++ {
		if haystack[i:i+l2] == needle {
			return i
		}
	}
	return -1
}
