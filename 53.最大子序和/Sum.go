package _3_最大子序和

func MaxSubArray(nums []int) int {
	var index, sum int
	var max=nums[0]
	var flag = true
	for _, v := range nums {
		index = index + v
		if index > sum {
			sum = index
		} else if index<0 {
			index = 0
		}
		if v>0 {
			flag=false
		}else {//全负数数组，找出最大的单个负数即可
			if v>max {
				max=v
			}
		}
	}
	if flag {
		return max
	}
	return sum
}
