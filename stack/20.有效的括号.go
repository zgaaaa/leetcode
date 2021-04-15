package stack

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

示例 2：
输入：s = "()[]{}"
输出：true

示例 3：
输入：s = "(]"
输出：false

示例 4：
输入：s = "([)]"
输出：false
*/

func IsValid(s string) bool {
	stack := []rune{}
	for _, v := range s {
		// 左括号入栈
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
			// 当遍历到右括号时,如果栈里没有元素返回false
		} else {
			if len(stack) == 0 {
				return false
			}
			// 栈顶元素
			top := stack[len(stack)-1]
			// 如果栈顶元素和遍历到的右括号相匹配,栈顶出栈,否则返回false
			if (v == ')' && top == '(') || (v == '}' && top == '{') || (v == ']' && top == '[') {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0 // 栈空则表名，匹配成功
}
