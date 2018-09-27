package _44_二叉树的前序遍历

import "testing"

func TestPreorderTraversal(t *testing.T) {
	var root = &TreeNode{Val:1,Left:&TreeNode{Val:1,Right:&TreeNode{Val:2}}}
	println(PreorderTraversal(root)) //[1,1,2]
}
