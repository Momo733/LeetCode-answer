package 二叉树的中序遍历

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
func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	
	res := InorderTraversal(root.Left)
	res = append(res, root.Val)
	res = append(res, InorderTraversal(root.Right)...)
	
	return res
}
