package __回文数

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	y:=x
	data := 0
	for y > 0 {
		data = data*10 + y%10
		y = y / 10
	}
	if data==x {
		return true
	}
	return false
}
