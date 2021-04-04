package binarytree

/*
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

示例：
二叉树：[3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
返回其层序遍历结果：
[
  [3],
  [9,20],
  [15,7]
]
*/

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func LevelOrder(root *TreeNode) [][]int {
	res := [][]int{} // 此时res为空 []
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	level := 0
	for len(queue) != 0 {
		// 提前给res添加一个空数组,此时res的状态为[[]]
		res = append(res, []int{})
		size := len(queue)
		for i := 0; i < size; i++ {
			n := queue[0]
			// 如果不提前添加一个空数组,这一步会越界
			res[level] = append(res[level], n.Val)
			queue = queue[1:]
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
		level++
	}
	return res
}
