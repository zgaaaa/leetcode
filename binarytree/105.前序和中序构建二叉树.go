package binarytree

/*
根据一棵树的前序遍历与中序遍历构造二叉树。
注意:
你可以假设树中没有重复的元素。

例如，给出
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：
    3
   / \
  9  20
    /  \
   15   7
*/

// 递归法:
func BuildTree(preorder []int, inorder []int) *TreeNode {
	// 当数组长度为零的时候跳出递归
	if len(preorder) == 0 {
		return nil
	}
	// 前序遍历的第一个节点就是根节点
	root := &TreeNode{Val: preorder[0]}
	// 用于计算中序遍历中根节点的位置
	i := 0
	// for ; i < len(inorder); i++ {
	// 	if inorder[i] == root.Val {
	// 		break
	// 	}
	// }
	for ; inorder[i] != root.Val; i++ {
	}
	// 递归求左右子节点
	root.Left = BuildTree(preorder[1:i+1], inorder[:i])
	root.Right = BuildTree(preorder[i+1:], inorder[i+1:])
	return root
}
