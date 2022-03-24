package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{}, 6))
}

func removeElement(nums []int, val int) int {

	index :=0
	for i:=0;i<len(nums);i++ {
		if nums[i]!=val {
			nums[index] = nums[i]
			index++
		}
	}

	return index
}