package _45_二叉树的后序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PostorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	
	res := PostorderTraversal(root.Left)
	res = append(res, PostorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res
}