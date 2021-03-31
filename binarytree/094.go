package binarytree

/*
中序遍历
*/

func InorderTraversal(root *TreeNode) []int {
	res := []int{}
	// 递归栈
	stark := []*TreeNode{}
	// 最左边的节点一个个压入递归栈
	for root != nil {
		stark = append(stark, root)
		root = root.Left
	}
	// 
	for len(stark) != 0 {
		// 栈顶位置
		top := len(stark) - 1
		// 输出栈顶节点
		res = append(res, stark[top].Val)
		// 栈顶(树的底部)节点的右子节点入栈
		root = stark[top].Right
		// 栈顶出栈
		stark = stark[:top]
		// 新入栈的右子节点的左节点一个个入栈,和上面一样
		for root != nil {
			stark = append(stark, root)
			root = root.Left
		}
	}
	return res
}
