package _3_搜索旋转排序数组

func Search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left<=right  {
		mid:=int((left+right)/2)
		if nums[mid]==target {
			return mid
		}
		if nums[mid]>=nums[left] { //左侧有序
			if nums[left]<=target&&target<nums[mid] { //目标值处在这个左侧有序数组内
				right=mid-1
			}else {
				left=mid+1
			}
		}else {//右侧有序
			if nums[mid]<=target&&target<=nums[right] {//目标值处在这个右侧有序数组内
				left=mid
			}else {
				right=mid-1
			}
		}
	}
	return -1
}
