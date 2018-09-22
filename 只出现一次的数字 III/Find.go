package 只出现一次的数字_III

func SingleNumber(nums []int) []int {
	var xor int
	for _, num := range nums {
		xor ^= num
	}
	//println(xor)  4^6
	lowest := xor & -xor //(4^6)& -(4^6) & 4
	//println(lowest) 2
	var a, b int
	for _, num := range nums {
		if num&lowest == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	//println(a,b) 4,6
	return []int{a, b}
}
