package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates([]int{1,1,2,2,3}))
}

func removeDuplicates(nums []int) int {
	a:=0
	for i:=0;i<len(nums)-1;i++ {
		if nums[i]!=nums[i+1] {
			nums[a] = nums[i]
			a++
			nums[a] = nums[i+1]
		}
	}
	nums = nums[0:a+1]
	return len(nums)
}