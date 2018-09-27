package _62_寻找峰值

func FindPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1
	if len(nums) == 1||nums[0] > nums[1] { //只有一个元素和第一个元素是峰值的时候
		return 0
	}
	if nums[len(nums)-1] > nums[len(nums)-2] {//最后的一个元素是峰值的时候
		return len(nums) - 1
	}
	for left < right { //中间的元素是峰值的时候
		mid := int((right + left) / 2)
		if nums[mid] > nums[mid+1] && nums[mid] > nums[mid-1] {
			return mid
		}
		if nums[mid] < nums[mid+1] {//右侧有峰值
			left = mid + 1
		} else { //左侧有峰值
			right = mid
		}
	}
	return -1
}
