package _55_最小栈

import "testing"

func TestMinStack_GetMin(t *testing.T) {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	println(stack.GetMin())
	stack.Pop()
	println(stack.Top())
	println(stack.GetMin())
}