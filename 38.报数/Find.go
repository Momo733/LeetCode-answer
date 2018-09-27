package _8_报数

import "strconv"

func CountAndSay(n int) string {
	var res = "1"
	for i := 1; i < n; i++ {
		res = say(res)
	}
	
	return res
}

func say(s string) string {
	des := s[:1]
	count := 0
	res := ""
	for i, k := range s[:] {
		if string(k) == des {
			count++
		} else {
			res = res + strconv.Itoa(count) + des
			des = string(k)
			count = 1
		}
		if i == len(s)-1 {
			res = res + strconv.Itoa(count) + des
		}
	}
	return res
}
