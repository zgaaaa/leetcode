package binarytree

/*
给定一个二叉树，判断其是否是一个有效的二叉搜索树。
假设一个二叉搜索树具有如下特征：
节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:
输入:
    2
   / \
  1   3
输出: true

示例 2:
输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。
*/

// 方法一：BFS（queue）
func IsValidBST(root *TreeNode) bool {
	queue := []*TreeNode{root}
	maxque := []int{1<<63 - 1}
	minque := []int{-1 << 63}
	for len(queue) != 0 {
		root, max, min := queue[0], maxque[0], minque[0]
		if root.Val <= min || root.Val >= max {
			return false
		}
		queue, maxque, minque = queue[1:], maxque[1:], minque[1:]
		if root.Left != nil {
			queue = append(queue, root.Left)
			maxque = append(maxque, root.Val)
			minque = append(minque, min)
		}
		if root.Right != nil {
			queue = append(queue, root.Right)
			maxque = append(maxque, max)
			minque = append(minque, root.Val)
		}
	}
	return true
}

// 方法二:中序遍历
// 二叉查找树的中序遍历结果是升序的
func IsValidBST2(root *TreeNode) bool {
	min := -1 << 63
	// 递归栈
	stark := []*TreeNode{}
	for root != nil || len(stark) != 0 {
		// 先将左子树一个个压入栈
		for root != nil {
			stark = append(stark, root)
			root = root.Left
		}
		// 栈顶
		top := len(stark) - 1
		// 如果栈顶节点的值小与下界返回false
		if stark[top].Val <= min {
			return false
		}
		// 下界变为栈顶节点的值
		min = stark[top].Val
		// 栈顶(即树的最底部)的右子节点标记为root
		root = stark[top].Right
		// 栈顶节点出栈
		stark = stark[:top]	
	}
	return true
}
