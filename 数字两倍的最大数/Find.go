package 数字两倍的最大数

func DominantIndex(nums []int) int {
	var last, second, lastIndex int
	for i, v := range nums {
		if v >= last {
			second = last
			lastIndex = i
			last = v
		}
		if v < last && v > second {
			second = v
		}
	}
	if last/2 < second {
		return -1
	} else {
		return lastIndex
	}
}
