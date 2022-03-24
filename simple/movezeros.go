package main

import "fmt"

//NO.283
func main() {
	moveZeroes([]int{1,0,1,0,888,0,9,0})
}

//快慢指针
func moveZeroes(nums []int)  {
	l:=0 //左指针
	r:=0 //右指针
	for r<len(nums) {
		if nums[r] !=0 {
			m:= nums[r]
			nums[r] = nums[l]
			nums[l] = m
			l++
		}
		r++
	}
	fmt.Println(nums)
}
