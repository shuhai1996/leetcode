package main

import "fmt"

//NO.6020
func main() {
	fmt.Println(divideArray([]int{1}))
}

func divideArray(nums []int) bool {
	m:= make(map[int]int)
	for i:=0;i< len(nums);i++ {
		if _,ok := m[nums[i]];ok {
			m[nums[i]] += 1
		} else {
			m[nums[i]] =1
		}

		if len(m) > len(nums)/2 {
			return false
		}
	}

	for _,n:=range m{
		if n%2 !=0 {
			return false
		}
	}
	return true
}