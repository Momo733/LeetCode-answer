package _3_罗马数字转整数

func RomanToInt(s string) int {
	var sMap = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
		//"IV": 4,
		//"IX": 9,
		//"XL": 40,
		//"XC": 90,
		//"CD": 400,
		//"CM": 900,
	}
	var target int
	var d int
	var last int //上一个参数的值
	//III
	for i := len(s) - 1; i >= 0; i-- {
		if sMap[s[i:i+1]] >= last {
			d = sMap[s[i:i+1]]
		} else {
			d = -sMap[s[i:i+1]]
		}
		last = sMap[s[i:i+1]]
		target = target + d
	}
	
	return target
}
