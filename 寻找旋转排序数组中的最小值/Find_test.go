package 寻找旋转排序数组中的最小值

import "testing"

func TestFindMin(t *testing.T) {
	var nums = []int{1, 2}
	println(FindMin(nums)) //1
	nums = []int{5,4,1, 2}
	println(FindMin(nums)) //1
	nums = []int{3,1, 2}
	println(FindMin(nums)) //1
	nums = []int{2,1}
	println(FindMin(nums)) //1
	nums = []int{4,1, 2}
	println(FindMin(nums)) //1
}
