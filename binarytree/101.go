package binarytree

/*
给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
    1
   / \
  2   2
 / \ / \
3  4 4  3

但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
    1
   / \
  2   2
   \   \
   3    3
*/

//解题思路：根节点的左右子树对称

// 方法一：递归
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(l, r *TreeNode) bool {
	switch {
	case l == nil && r == nil:
		return true
	case l == nil || r == nil:
		return false
	case l.Val != r.Val:
		return false
	}
	return isMirror(l.Left, r.Right) && isMirror(l.Right, r.Left)
}

// 方法二：迭代
func IsSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) != 0 {
		l, r := queue[0], queue[1]
		queue = queue[2:]
		switch {
		case l == nil && r == nil :// 这个判断是为了不让下个case产生误判
			continue
		case l == nil || r == nil:
			return false
		case l.Val != r.Val:
			return false
		}
		queue = append(queue, l.Left, r.Right)
		queue = append(queue, l.Right, r.Left)
	}
	return true
}
