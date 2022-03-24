package main

import (
	"fmt"
)
//NO.704
func main() {
	fmt.Println(search([]int{-1,0,3,5,9,12}, 3))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l:=0
	h:=len(nums)-1
	for l<= h {
		half:= h-l/2 + l
		if nums[half] == target {
			return half
		}else if nums[half]<= target {
			l = half +1
		} else {
			h = half -1
		}
	}
	return -1
}

