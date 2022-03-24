package main

import (
	"fmt"
)

//NO.6027
func main() {

	fmt.Println(countHillValley([]int{5,7,5,7,5,6}))
}

func countHillValley(nums []int) int {
	d := 0
	oldPos, pos, count := 0,0,0

	for i:=1;i<len(nums);i++ {
		d = diff(nums[i],nums[i-1])
		if nums[i] >nums[i-1] {
			pos = 1
		} else if nums[i] <nums[i-1] {
			pos = -1
		}

		if pos == 0-oldPos && d>0{
			count++
		}
		oldPos = pos
	}
	return count
}

func diff(a int ,b int) int {
	if a > b {
		return a-b
	} else {
		return b-a
	}
}