package _69__求众数

func MajorityElement(nums []int) int {
	quickSort(nums, 0, len(nums)-1)
	return nums[int((len(nums)-1)/2)]
}

func quickSort(nums []int, left, right int) {
	temp := nums[left]
	p := left
	i, j := left, right
	
	for i <= j {
		for j >= p && nums[j] >= temp {
			j--
		}
		if j >= p {
			nums[p] = nums[j]
			p = j
		}
		
		if nums[i] <= temp && i <= p {
			i++
		}
		if i <= p {
			nums[p] = nums[i]
			p = i
		}
	}
	nums[p] = temp
	if p-left > 1 {
		quickSort(nums, left, p-1)
	}
	if right-p > 1 {
		quickSort(nums, p+1, right)
	}
}

