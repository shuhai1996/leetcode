package main

import (
	"fmt"
)

//NO.66
func main() {
	fmt.Println(plusOne([]int{7,2,8,5,0,9,1,2,9,5,3,6,6,7,3,2,8,4,3,7,9,5,7,7,4,7,4,9,4,7,0,1,1,1,7,4,0,0,6}))
}

func plusOne(digits []int) []int {
	l := len(digits)
	for i:=l-1;i>=0;i-- {
		if digits[i] != 9 {
			digits[i]++
			for j:=i+1;j<l; j++ {
				digits[j] =0
			}
			return digits
		}
	}
	digits = make([]int, l+1)
	digits[0] = 1
	return digits
}