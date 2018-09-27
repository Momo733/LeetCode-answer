package _62_寻找峰值

import "testing"

func TestFindPeakElement(t *testing.T) {
	var nums = []int{1}
	println(FindPeakElement(nums)) //0
	nums = []int{1, 2}
	println(FindPeakElement(nums)) //1
	nums = []int{2, 1}
	println(FindPeakElement(nums)) //0
	nums = []int{1, 3, 2}
	println(FindPeakElement(nums)) //1
}
