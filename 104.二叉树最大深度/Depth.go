package _04_二叉树最大深度
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func MaxDepth(root *TreeNode) int {
	if root==nil {
		return 0
	}
	
	return 1 + max(MaxDepth(root.Left), MaxDepth(root.Right))
}

func max(x,y int) int{
	if x>y {
		return x
	}
	return y
}