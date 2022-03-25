package main

import "fmt"
//191
func main() {
	fmt.Println(hammingWeight(9))
}

func hammingWeight(num uint32) int {
	result := 0
	for i:=0;i<32;i++ {
		if 1<<i&num >0 {
			result++
		}
	}
	return result
}