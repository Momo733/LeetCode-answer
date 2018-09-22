package 最大子序和

import "testing"

func TestMaxSubArray(t *testing.T) {
	var nums = []int{-2,-1,5}
	println(MaxSubArray(nums))
	nums=[]int{-1,-1}
	println(MaxSubArray(nums)) //-1
	nums=[]int{-1}
	println(MaxSubArray(nums)) //-1
}
