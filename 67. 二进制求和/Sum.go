package _7__二进制求和
func AddBinary(a string, b string) string {
	if len(b) < len(a) {
		for i := len(a) - len(b); i > 0; i-- {
			b = "0" + b
		}
	}
	if len(b) > len(a) {
		for i := len(b) - len(a); i > 0; i-- {
			a = "0" + a
		}
	}
	var index int
	var res = ""
	var j = len(b) - 1
	for i := len(a) - 1; (0 <= i && i < len(a)) || (0 <= j && j < len(b)); i-- {
		if index == 0 && a[i:i+1] == "0" && b[j:j+1] == "0" {
			index = 0
			res = "0" + res
			j--
			continue
		}
		if index == 0 && a[i:i+1] == "1" && b[j:j+1] == "1" {
			index = 1
			res = "0" + res
			j--
			continue
		}
		if index == 0 && (a[i:i+1] == "1" || b[j:j+1] == "1") {
			index = 0
			res = "1" + res
			j--
			continue
		}
		
		if index == 1 && a[i:i+1] == "0" && b[j:j+1] == "0" {
			index = 0
			res = "1" + res
			j--
			continue
		}
		if index == 1 && a[i:i+1] == "1" && b[j:j+1] == "1" {
			index = 1
			res = "1" + res
			j--
			continue
		}
		if index == 1 && (a[i:i+1] == "1" || b[j:j+1] == "1") {
			index = 1
			res = "0" + res
			j--
			continue
		}
	}
	if index == 1 {
		return "1"+res
	}
	return res
}
