package main

func main() {

}
//NO.20
func isValid(s string) bool {
	if len(s)%2 !=0 {
		return false
	}
	var stack []byte
	m := map[byte]byte{
		')':'(',
		']':'[',
		'}':'{',
	}
	for i:=0;i<len(s);i++ {
		if m[s[i]] > 0 {
			if len(stack)==0 || stack[len(stack)-1] != m[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1] //出栈
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack)==0 //全部出栈了说明都匹配上了，否则说明有没匹配上的
}
