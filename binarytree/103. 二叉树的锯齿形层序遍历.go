package binarytree

/*
给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层序遍历如下：

[
  [3],
  [20,9],
  [15,7]
]
*/

func ZigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	level := 0
	for len(queue) != 0 {
		size := len(queue)
		p := make([]int,size) // 因为每一层的个数是知道的,可以提前声明大小
		for i := 0; i < size; i++ {
			n := queue[0]
			queue = queue[1:]
			if level%2 == 0 {
				p[i] = n.Val // 利用变量 i
			} else {
				p[size-i-1] = n.Val
			}
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
		res = append(res, p)
		level++
	}
	return res
}
