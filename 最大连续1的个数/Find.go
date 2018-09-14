package 最大连续1的个数

func FindMaxConsecutiveOnes(nums []int) int {
	var count int
	var index int
	var flag bool
	for _, v := range nums {
		if v == 1 {
			flag = true
			index++
			if index > count {
				count = index
			}
		} else {
			if index > count {
				count = index
			}
			index = 0
		}
	}
	if !flag {
		count = 0
	}
	return count
}
