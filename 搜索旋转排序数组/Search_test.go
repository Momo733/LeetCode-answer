package 搜索旋转排序数组

import "testing"

func TestSearch(t *testing.T) {
	var nums = []int{4, 5, 6, 7, 0, 1, 2}
	var target = 3
	println(Search(nums, target))
	nums = []int{4, 5, 6, 7, 0, 1, 2}
	target = 0
	println(Search(nums,target))
}
