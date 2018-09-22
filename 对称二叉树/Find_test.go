package 对称二叉树

import "testing"

func TestIsSymmetric(t *testing.T) {
		var root =&TreeNode{Val:1,Left:&TreeNode{Val:2,Left:&TreeNode{Val:3},Right:&TreeNode{Val:3}},Right:&TreeNode{Val:2,Left:&TreeNode{Val:3},Right:&TreeNode{Val:3}}}
		println(IsSymmetric(root)) //true
}
