package main

import "fmt"

func main() {
	fmt.Println(isPalindromeV2(988))
}

func isPalindromeV2(x int) bool {
	if x<0 {
		return false //0或负数直接返回false
	}
	if x==0 {
		return true //0直接返回true
	}
	n:=x
	a:= 0
	for n>0 {
		a = a*10+ n%10
		n = n/10
	}
	fmt.Println(a)
	return a==x
}

/* 是回文数*/
func isPalindromeV1(x int) bool {
	var strArr []int
	if x<0 {
		return false //0或负数直接返回false
	}
	if x==0 {
		return true //0直接返回true
	}
	for x>0 {
		strArr = append(strArr, x%10)
		x = x/10
	}
	i:=0
	for i<=len(strArr)/2 {
		if strArr[i] != strArr[len(strArr)-1-i] {
			return false
		}
		i++
	}
	return true
}