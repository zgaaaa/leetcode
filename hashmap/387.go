package hashmap

/*
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

示例：
s = "leetcode"
返回 0

s = "loveleetcode"
返回 2

思路与算法:
我们可以对字符串进行两次遍历。
在第一次遍历时，我们使用哈希映射统计出字符串中每个字符出现的次数。在第二次遍历时，
我们只要遍历到了一个只出现一次的字符，那么就返回它的索引，否则在遍历结束后返回−1。
*/

func FirstUniqChar(s string) int {
	//创建一个数组
	arr := [26]int{}
	// 遍历字符串,统计字符个数
	for _, v := range s {
		arr[v-'a']++
	}
	// 题目要求是返回原字符串中的第一个不重复字符,所以需要遍历字符串s
	for i, v := range s {
		if arr[v-'a'] == 1 {
			return i
		}
	}
	// 这里的错误做法
	//这样做返回的是数组arr中的第一个不重复字符的下标,不是原字符串s的
	// for i, v := range arr {
	// 	if v == 1 {
	// 		return i
	// 	}
	// }
	return -1
}
