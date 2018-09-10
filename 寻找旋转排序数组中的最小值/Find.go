package 寻找旋转排序数组中的最小值

func FindMin(nums []int) int {
	low := 0
	high := len(nums) - 1
	target := nums[0]
	if nums[low] <= nums[high] { //旋转之后和原来的数组相同的情况下,第一个值是最小值
		return nums[low]
	}
	if nums[low] > nums[high] && nums[high] < nums[high-1] { //完全倒叙的情况
		return nums[high]
	}
	for low < high {
		mid := int((high + low) / 2)
		if nums[mid] < nums[mid+1] && nums[mid] < nums[mid-1] { //寻找最低点,判断该中心点事最低点直接返回
			return nums[mid]
		}
		if nums[mid] >= nums[low] { //左侧是升序，说明最小点在右侧
			low = mid + 1
			if target >= nums[mid+1] {
				target = nums[mid+1]
			}
		} else { //右侧是升序，说明最小点在左侧
			high = mid - 1
			if target >= nums[mid-1] {
				target = nums[mid-1]
			}
		}
	}
	return target
}
