package binarytree

/*
给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

示例 1：
输入：p = [1,2,3], q = [1,2,3]
输出：true

示例 2：
输入：p = [1,2], q = [1,null,2]
输出：false

示例 3：
输入：p = [1,2,1], q = [1,1,2]
输出：false
*/

// 方法一：迭代
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	queue := []*TreeNode{p, q}
	for len(queue) > 0 {
		l, r := queue[0], queue[1]
		queue = queue[2:]
		switch {
		case l == r:
			continue
		case l == nil || r == nil:
			return false
		case l.Val != r.Val:
			return false
		}
		queue = append(queue, l.Left, r.Left)
		queue = append(queue, l.Right, r.Right)
	}
	return true
}

// 方法二:递归
func IsSameTree2(p *TreeNode, q *TreeNode) bool {
	switch {
	case p == nil && q == nil:
		return true
	case p == nil || q == nil:
		return false
	case p.Val != q.Val:
		return false
	}
	return IsSameTree2(p.Left, q.Left) && IsSameTree2(p.Right, q.Right)
}
