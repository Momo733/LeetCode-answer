package _44__反转字符串

func  ReverseString(s string) string {
	var i = len(s)-1
	var res = []byte(s)
	for j := 0; j < i; j++ {
		res[i],res[j]=res[j],res[i]
		i--
	}
	return string(res)
}