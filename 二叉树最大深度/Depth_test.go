package 二叉树最大深度

import "testing"

func TestMaxDepth(t *testing.T) {
	var root =&TreeNode{Val:1,Left:&TreeNode{Val:2,Right:&TreeNode{Val:3}},Right:&TreeNode{Val:5,Right:&TreeNode{Val:9}}}
	println(MaxDepth(root))
}
