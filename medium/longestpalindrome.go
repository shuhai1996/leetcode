package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome(""))
}

func longestPalindrome(s string) string {
	var position int
	strArr := []byte(s)
	var currentStrArr []byte // 当前字符串数组
	for i:=len(strArr);i>0;i-- {
		position = 0
		for i+position <= len(strArr){
			currentStrArr = strArr[position:i+position]
			if isCircleSt(currentStrArr) {
				return string(currentStrArr)
			}
			position++
		}
	}
	return ""
}

/* 是回文串*/
func isCircleSt(strArr []byte) bool {
	i:=0
	for i<=len(strArr)/2 {
		if strArr[i] != strArr[len(strArr)-1-i] {
			return false
		}
		i++
	}
	return true
}