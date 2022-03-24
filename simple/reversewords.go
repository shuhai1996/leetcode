package main
//NO.557
import (
	"fmt"
)

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

func reverseWords(s string) string {
	b:= []byte(s)
	var newB []byte
	pos:=0
	for i:=0;i<len(b);i++ {
		if b[i] ==' ' {
			cur:= reverseS(b[pos:i])
			newB = append(newB, cur...)
			if b[i] ==' ' {
				newB = append(newB, ' ')
			}
			pos =i+1
		} else if i==len(b)-1 {
			cur:= reverseS(b[pos:i+1])
			newB = append(newB, cur...)
		}
	}
	return string(newB)
}

func reverseS(s []byte) []byte {
	var m byte
	l := len(s)
	for i:=0;i<(len(s)+1)/2;i++ {
		m = s[i]
		s[i] = s[l-1-i]
		s[l-1-i] = m
	}
	return s
}