package 只出现一次的数字_II

import "sort"

func SingleNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 3 {
		if len(nums)-1 == i {
			return nums[i]
		}
		if nums[i]^nums[i+1] != 0 && nums[i+1]^nums[i+2] != 0 {
			return nums[i+1]
		}
		if nums[i+1]^nums[i+2] == 0 && nums[i]^nums[i+1] != 0 {
			return nums[i]
		}
		if nums[i+1]^nums[i+2] != 0 && nums[i]^nums[i+1] == 0 {
			return nums[i+2]
		}
	}
	return 0
}