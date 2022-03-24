package main
//NO.567
import (
	"fmt"
)

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
}

func checkInclusion(s1 string, s2 string) bool {
	n:= len(s1)
	m:= len(s2)
	if n > m {
		return false
	}
	var str1, str2 [26]int
	for i,ch := range s1{
		str1[ch-'a'] ++
		str2[s2[i]-'a']++
	}
	if str1 == str2 {
		return true
	}

	for i:=n; i<m ;i++ {
		str2[s2[i]-'a']++
		str2[s2[i-n]-'a']--
		if str1 == str2 {
			return true
		}
	}
	return false
}

//反转字符串
func reverseStr(s string) string {
	str := []byte(s)
	var m byte
	l := len(s)
	for i:=0;i<(len(str)+1)/2;i++ {
		m = str[i]
		str[i] = s[l-1-i]
		str[l-1-i] = m
	}
	return string(str)
}