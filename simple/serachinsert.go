package main

import "fmt"

//NO.35
func main() {
	fmt.Println(searchInsert([]int{1,3,5,6}, 8))
}

func searchInsert(nums []int, target int) int {
	l:= 0
	h:= len(nums)-1
	for l<=h {
		half:= (l+h)/2
		if nums[half] == target {
			return half
		}else if nums[half] < target {
			l = half+1
		}else {
			h = half-1
		}
	}
	return l
}