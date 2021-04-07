package binarytree

/*
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
*/

func IsBalanced(root *TreeNode) bool {
	temp := true
	height(root, &temp)
	return temp
}
// 递归函数
func height(root *TreeNode, temp *bool) int {
	if root == nil {
		return 0
	}
	l := height(root.Left, temp)
	r := height(root.Right, temp)
    // 如果左右子树的深度的绝对值大于1,则temp为false
	if abs(l, r) > 1 {
		*temp = false
	}
    // 返回深度
	return max(l, r) + 1
}
func abs(a, b int) int {
	if a-b < 0 {
		return b - a
	}
	return a - b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
