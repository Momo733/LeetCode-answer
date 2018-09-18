package 有效的括号

func IsValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}
	var sMap =map[rune]rune{
		')':'(',
		'}':'{',
		']':'[',
	}
	st := new(stack)
	for _, v := range s[:] {
		if string(v) == "(" || string(v) == "[" || string(v) == "{" {
			st.push(v)
		} else {
			if data, success := st.pop(); !success ||data!=sMap[v] {
				return false
			}
		}
	}
	if len(*st)>0 {
		return false
	}
	return true
}

type stack []rune

func (s *stack) push(b rune) {
	*s = append(*s, b)
}
func (s *stack) pop() (data rune, success bool) {
	if len(*s) > 0 {
		data = (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return data, true
	}
	return 0, false
}

