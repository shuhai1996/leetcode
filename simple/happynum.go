package main

import (
	"fmt"
)
func main() {
	//var a int
	//i :=0
	//n:=19
	//for n>0 {
	//	a += (n%10) * (n%10)
	//	n = n/10
	//	i++
	//}
	//fmt.Println(a)
	//fmt.Println(i)
	//os.Exit(1)
	fmt.Println(isHappy(2))
}

func isHappy(n int) bool {
	var a int
	i :=0
	var noHappy = []int{4,16,37,58,89,145,42,20}//不快乐数
	for _,i := range noHappy{
		if n == i {
			return false
		}
	}
	for n>0 {
		a += (n%10) * (n%10)
		n = n/10
		i++
	}
	fmt.Println(a)
	if a == 1 {
		return true
	} else if i>=1{
		n = a
		return isHappy(n)
	}
	return false
}