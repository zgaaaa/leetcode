package dfs

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例 1：
输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

示例 2：
输入：digits = ""
输出：[]

示例 3：
输入：digits = "2"
输出：["a","b","c"]
*/

func LetterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	res := []string{""}
	m := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	// 遍历每个数字
	for i := 0; i < len(digits); i++ {
		// 每个数字所映射的字母
		letter := m[digits[i]-'2']
		size := len(res)
		// 从队首弹出一个字符串与letter里的字符串拼接,然后添加到队尾
		for j := 0; j < size; j++ {
			// 出队
			temp := res[0]
			res = res[1:]
			// 字符串拼接
			for k := 0; k < len(letter); k++ {
				res = append(res, temp+string(letter[k]))
			}
		}
	}
	return res
}
