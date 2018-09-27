package _36_只出现一次的数字
func SingleNumber(nums []int) int {
	var res int
	for _,v:=range nums{
		res^=v
	}
	return res
}
