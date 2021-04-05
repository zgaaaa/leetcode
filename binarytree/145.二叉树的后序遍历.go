package binarytree

/*
后序遍历
*/

func PostorderTraversal(root *TreeNode) []int {
	// 返回结果
	res := []int{}
	if root == nil {
		return res
	}
	// 栈
	stark := []*TreeNode{root}
	// 存储节点
	mapnode := make(map[*TreeNode]int)
	for len(stark) != 0 {
		// 出栈
		top := len(stark) - 1
		root = stark[top]
		stark = stark[:top]
		// 如果节点已经被遍历过,直接打印
		if _, ok := mapnode[root]; ok {
			res = append(res, root.Val)
		// 如果是第一次出现,则将当前节点,右子节点及左子节点入栈
		}else {
			stark = append(stark, root)
			if root.Right != nil {
				stark = append(stark, root.Right)
			}
			if root.Left != nil {
				stark = append(stark, root.Left)
			}
			// 将当前节点加入map
			mapnode[root] = 1
		}
	}
	return res
}

// 递归解法:
func PostorderTraversal2(root *TreeNode) (res []int) {
	var recursion func(*TreeNode)
	recursion = func(root *TreeNode) {
		if root == nil {
			return
		}
		recursion(root.Left)
		recursion(root.Right)
		res = append(res, root.Val)
	}
	recursion(root)
	return 
}