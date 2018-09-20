package 二叉树的后序遍历

import "testing"

func TestPostorderTraversal(t *testing.T) {
	var root = &TreeNode{Val:1,Left:&TreeNode{Val:1,Right:&TreeNode{Val:2}}}
	println(PostorderTraversal(root))
}
