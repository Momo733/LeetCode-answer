package 二叉树的中序遍历

import "testing"

func TestInorderTraversal(t *testing.T) {
	var root = &TreeNode{Val:1,Left:&TreeNode{Val:1,Right:&TreeNode{Val:2}}}
	println(InorderTraversal(root))
}
