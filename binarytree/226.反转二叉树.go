package binarytree

/*
反转二叉树
*/

// 方法一:迭代
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		top := queue[0]
		top.Left, top.Right = top.Right, top.Left
		queue = queue[1:]
		if top.Left != nil {
			queue = append(queue, top.Left)
		}
		if top.Right != nil {
			queue = append(queue, top.Right)
		}
	}
	return root
}

// 方法二:递归
func InvertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	InvertTree2(root.Left)
	InvertTree2(root.Right)
	return root
}
