package 对称二叉树


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func IsSymmetric(root *TreeNode) bool {
	if root==nil{
		return true
	}
	
	return Symmetric(root.Left,root.Right)
}

func Symmetric(left,right *TreeNode) bool{
	if left==nil&&right==nil {
		return true
	}
	if left==nil||right==nil {
		return false
	}
	return left.Val==right.Val&&Symmetric(left.Left,right.Right)&&Symmetric(left.Right,right.Left)
}
