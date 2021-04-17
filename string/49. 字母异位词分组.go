package string

import "sort"

/*
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:
输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
说明：

所有输入均为小写字母。
不考虑答案输出的顺序。
*/

func GroupAnagrams(strs []string) [][]string {
	res := [][]string{}
	m := make(map[string][]string, len(strs))
	for _, v := range strs {
		s := []byte(v)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] }) // 排序
		key := string(s) // 异位词排序后一定是相同的
		// map的键是排序后的字符串,值是字符串切片,用于存储分组
		m[key] = append(m[key], v)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
