/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stark := []*TreeNode{root} // 根节点先入栈
	for len(stark) > 0 {
		top := stark[len(stark)-1] // 栈顶节点
		res = append(res, top.Val) // 输出
		stark = stark[:len(stark)-1] // 栈顶出栈
		// 因为栈是先进后出,所以右节点先进栈
		if top.Right != nil {
			stark = append(stark, top.Right)
		}
		if top.Left != nil {
			stark = append(stark, top.Left)
		}
	}
	return res
}
// @lc code=end

