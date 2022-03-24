package main
//NO.977
import (
	"fmt"
)

func main() {
	fmt.Println(sortedSquares([]int{-4,-1,0,3,10}))
}

//暴力解法，冒泡排序
func sortedSquaresV1(nums []int) []int {
	//var m = make([]int, len(nums))
	for index,i:=range nums{
		nums[index] = i*i
	}

	for i:=0;i<len(nums);i++ {
		for j:=i;j<len(nums);j++ {
			if nums[i]> nums[j] {
				k:= nums[i]
				nums[i] = nums[j]
				nums[j] = k
			}
		}
	}

	return nums
}

//双指针 , 解法一
func sortedSquaresV2(nums []int) []int {
	var m = make([]int, len(nums)) //创建一个长度为目标长度的整型数组
	l:=0 //左指针
	r:=len(nums)-1 //右指针
	for position:=len(nums)-1;position>=0;position-- { //指针划到最后
		j:= nums[l] * nums[l]
		k:= nums[r] * nums[r]
		if j>k {
			m[position] = j
			l++
		} else {
			m[position] = k
			r--
		}
	}
	return m
}

//双指针 , 解法二,注意边界问题
func sortedSquares(nums []int) []int {
	var m = make([]int, len(nums)) //创建一个长度为目标长度的整型数组
	var mid int //mid为边界将数组拆成两个有序数组，一个单调递减，一个单调递增
	for i:=0;i<len(nums) && nums[i]<0;i++ {
		mid = i
	}
	l:= mid //左指针从mid开始
	r:= mid+1 //右指针从mid+1开始
	for position:=0;position<len(nums);position++ {
		if l< 0 {
			m[position] = nums[r] *nums[r]
			r++
		}else if r>len(nums)-1 {
			m[position] = nums[l] *nums[l]
			l--
		}else if nums[l]*nums[l] < nums[r] * nums[r] {
			m[position] = nums[l] *nums[l]
			l--
		} else {
			m[position] = nums[r] *nums[r]
			r++
		}
	}
	return m
}

