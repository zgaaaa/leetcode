package binarytree

/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，
满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

示例 1：
输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出：3
解释：节点 5 和节点 1 的最近公共祖先是节点3。

示例 2：
输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出：5
解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。

示例 3：
输入：root = [1,2], p = 1, q = 2
输出：1
*/

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 如果当前节点是空,则不存在左右子节点
	if root == nil {
		return nil
	}
	// 如果当前节点为p或q,有两种可能
	// 1.祖先就是当前节点
	// 2.祖先在当前节点的上面
	// 所以直接返回root
	if root == p || root == q {
		return root
	}
	// 当前节点既不是空也不是pq,那就向下递归
	l := LowestCommonAncestor(root.Left, p, q)
	r := LowestCommonAncestor(root.Right, p, q)
	
	switch {
		// 如果l和r都有结果,说明p和q在root两侧,返回root
	case l != nil && r != nil :
		// 如果只有一个子树有结果,说明另一个节点不在这棵树上,返回结果
		return root
	case l == nil :
		return r
	case r == nil :
		return l
		// 如果都为nil,说明p和q不在这颗之数中,返回nil
	default:
		return nil
	}
}
