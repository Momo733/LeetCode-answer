package 回文链表

import "testing"

func TestIsPalindrome(t *testing.T) {
	var head = &ListNode{Val:1,Next:&ListNode{Val:2,Next:&ListNode{Val:1,Next:nil}}}
	println(IsPalindrome(head)) //true
}