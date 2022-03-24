package main

import "fmt"

func main() {
	fmt.Println(calculate( "(1+(4+5+2)-3)+(6+8)"))
}
//NO.224
//括号展开 + 栈
func calculate(s string) int {
	options := []int{1} //操作符
	sign := 1
	result := 0
	for i:=0;i<len(s); {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = options[len(options)-1]
			i++
		case '-':
			sign = -options[len(options)-1]
			i++
		case '(': //碰到(符号时进栈
			options = append(options, sign)
			i++
		case ')': //碰到)符号时出栈
			options = options[:len(options)-1]
			i++
		default:
			num:=0
			for ;i<len(s) && s[i] >='0' && s[i]<= '9';i++ {
				num = num*10 + int(s[i] - '0')
			}
			result += num*sign
		}
	}
	return result
}

