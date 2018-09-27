package _44_二叉树的前序遍历

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
 //自己创建栈来存储节点
func PreorderTraversal(root *TreeNode) []int {
	var stack = []*TreeNode{}
	var res =[]int{}
	for cur:=root;root!=nil; {
		res=append(res,cur.Val)
		if cur.Left!=nil {
			if cur.Right!=nil {
				stack=append(stack,cur.Right)
			}
			cur=cur.Left
		}else {
			if cur.Right!=nil {
				cur=cur.Right
			}else {
				if len(stack)==0 {
					break
				}else {
					cur=stack[len(stack)-1]
					stack=stack[:len(stack)-1]
				}
			}
			
		}
		
	}
	return res
}
//利用函数的栈来实现存储节点
//func PreorderTraversalStack(root *TreeNode) []int  {
//	if root == nil {
//		return nil
//	}
//
//	if root.Left==nil&&root.Right==nil {
//		return []int{root.Val}
//	}
//
//	res:=PreorderTraversalStack(root.Left)
//	res=append(res,root.Val)
//	res=append(res,PreorderTraversalStack(root.Right)...)
//	return res
//}

