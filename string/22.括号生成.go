package string

/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]

示例 2：
输入：n = 1
输出：["()"]
*/

func GenerateParenthesis(n int) []string {
	res := []string{}
	var dfs func(l, r int, temp string)
	dfs = func(l, r int, temp string) {
		// 只要左括号小于给定的n就添加左括号
		if l < n {
			// 左括号数量加一,同时添加左括号
			dfs(l+1, r, temp+"(")
		}
		// 只有当右括号的数量小于左括号时才添加右括号
		if r < l {
			// 添加右括号,右括号数量加一
			dfs(l, r+1, temp+")")
		}
		// 有括号的数量等于n的时候返回
		if r == n {
			res = append(res, temp)
			return
		}
	}
	dfs(0, 0, "")
	return res
}
