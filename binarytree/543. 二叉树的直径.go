package binarytree

/*
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。

示例 :
给定二叉树

          1
         / \
        2   3
       / \
      4   5
返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。

注意：两结点之间的路径长度是以它们之间边的数目表示。
*/

func DiameterOfBinaryTree(root *TreeNode) int {
	res := 0
	dfs(root, &res)
	return res
}
// 递归函数
func dfs(root *TreeNode, res *int) int {
	// 空节点返回0,则叶子节点为1
	if root == nil {
		return 0
	}
	// 递
	l := dfs(root.Left, res)
	r := dfs(root.Right, res)
	// 直径取左右子树深度相加的最大值
	*res = maxv(l+r, *res)
	// 返回深度 + 1
	return maxv(l, r) + 1 
}
func maxv(a, b int) int {
	if a > b {
		return a
	}
	return b
}