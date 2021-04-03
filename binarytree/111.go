package binarytree
/*
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明：叶子节点是指没有子节点的节点。
*/

func MinDepth(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			top := queue[0]
			queue = queue[1:]
			if top.Left == nil && top.Right == nil {
				return depth + 1
			}
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
        depth++
	}
	return depth
}