package binarytree

/*
给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，
    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 数组模拟队列
	var queue []*TreeNode
	// 直接将根节点放入队列
	queue = append(queue,root)
	// depth深度
	depth := 0
	// 当队列不为空时循环以下操作
	for len(queue) > 0 {
		// size用于记录当前层节点的个数
		size := len(queue)
		// 遍历当前层中的节点
		for i := 0; i < size; i++ {
			// 获得队首元素,并且队首元素出队
			n := queue[0]
			queue = queue[1:]
			// 如果有左节点或右节点,则进队
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
		// 遍历完一层后深度就加一
		depth++
	}
	return depth
}
