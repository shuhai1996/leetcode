package main

import "fmt"

//NO.189
func main() {
	rotate([]int{-1,1,2}, 2)
}

func rotateV1(nums []int, k int)  {
	m:=make([]int, len(nums))
	board :=len(nums)-1 // 右边界
	l :=0 //左指针
	r :=len(nums)-1 //右指针
	for i:=0;i<k;i++ {
		r++
		l++
		if r>board {
			r=0
		}
		if l >board {
			l=0
		}
	}

	for position:=0;position<=board;position++ {
		if position<l {
			m[position] = nums[board+position-r]
		} else {
			m[position] = nums[position-l]
		}
	}
	for i:=0;i<board;i++ {
		nums[i] = m[i]
	}
}

//解法二。反转数组
func rotate(nums []int, k int)  {
	l :=len(nums)
	k = k%l
	reverse(nums)
	reverse(nums[0:k])
	reverse(nums[k:l])
}

func reverse(nums []int) []int {
	board := len(nums)-1
	for i := 0; i < (board+1)/2; i++ {
		m := nums[i]
		nums[i] = nums[board-i]
		nums[board-i] = m
	}
	return nums
}
