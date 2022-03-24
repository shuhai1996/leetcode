package main

import (
	"fmt"
	"math"
)

//NO.11
func main() {
	fmt.Println(maxArea([]int{1,2,3,4}))
}

func maxArea(height []int) int {
	l:=0
	r:= len(height)-1
	max:=0
	for l<r {
		z := math.Min(float64(height[l]), float64(height[r]))
		area := int(z) * (r-l)
		if area >max {
			max = area
		}
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return max
}
