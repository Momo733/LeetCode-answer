package _2_整数转罗马数字
func IntToRoman(num int) string {
	var sMap = map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
		4:    "IV",
		9:    "IX",
		40:   "XL",
		90:   "XC",
		400:  "CD",
		900:  "CM",
	}
	var nums = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var s string
	for i := 0; i < len(nums) && num > 0; i++ {
		if num < nums[i] {
			continue
		}
		for num >= nums[i] {
			num = num - nums[i]
			s = s + sMap[nums[i]]
		}
	}
	return s
}
