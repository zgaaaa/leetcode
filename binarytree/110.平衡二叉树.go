package binarytree

import "math"

/*
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
*/

func IsBalanced(root *TreeNode) bool {
   return height(root) >= 0
}

func height(root *TreeNode) int{
    if nil==root{
        return 0
    }
    left:=height(root.Left)
    if left==-1 {
        return -1
    }
    right:=height(root.Right)
    if right == -1{
        return -1
    } 
    if math.Abs(float64(left)-float64(right))>1{
        return -1
    }else{
        return int(math.Max(float64(left),float64(right)))+1
    }
}