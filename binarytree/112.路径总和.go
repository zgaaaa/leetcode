package binarytree

/*
给你二叉树的根节点 root 和一个表示目标和的整数 targetSum ，判断该树中是否存在根节点到叶子节点的路径，
这条路径上所有节点值相加等于目标和 targetSum。叶子节点是指没有子节点的节点。
*/

// 方法一:递归法
// 问题拆解:将判断"从根节点到叶子节点是否存在节点值相加等于给定值的路径"拆解成
// "判断其子节点到叶子节点是否存在节点值相加等于给定值-当前节点值"
// 递归返回条件:1.当root为空时直接返回; 2.当root为叶子节点时判断值是否为递归更新后的Sum值.
func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil { // 归
		return targetSum == root.Val
	}
	l := HasPathSum(root.Left, targetSum-root.Val) // 向下递的时候更新Sum值
	r := HasPathSum(root.Right, targetSum-root.Val)
	return l || r
}

// 方法二:迭代法BFS
func HasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// 两个队列，用于储存节点和和路径
	queue := []*TreeNode{root}
	quesum := []int{root.Val}
	for len(queue) != 0 {
		// 出队
		n := queue[0]
		p := quesum[0]
		queue = queue[1:]
		quesum = quesum[1:]
		// 判断:当遍历到叶子节点时,如果p等于给定的值就返回真
		if n.Left == nil && n.Right == nil {
			if p == targetSum {
				return true
			}
			continue
		}
		// 入队
		if n.Left != nil {
			queue = append(queue, n.Left)
			quesum = append(quesum, n.Left.Val+p)
		}
		if n.Right != nil {
			queue = append(queue, n.Right)
			quesum = append(quesum, n.Right.Val+p)
		}
	}
	return false
}
