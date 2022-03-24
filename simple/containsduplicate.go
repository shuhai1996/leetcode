package main
//NO.217
import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1,2,2,3}))
}

func containsDuplicate(nums []int) bool {
	n:= make(map[int]int)
	for i:=0;i<len(nums);i++{
		if _,ok := n[nums[i]];ok {
			return true
		} else {
			n[nums[i]] = 1
		}
	}
	return false
}