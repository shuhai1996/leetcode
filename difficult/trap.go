package main
//NO.42
import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
}

//双指针解法
func trap(height []int) int {
	l:=0
	r:=len(height)-1
	lm, rm :=0, 0
	area := 0
	for l<r {
		lm = int(math.Max(float64(lm), float64(height[l])))
		rm = int(math.Max(float64(rm), float64(height[r])))
		if height[l] < height[r] {
			area += lm - height[l]
			l++
		} else {
			area += rm - height[r]
			r--
		}
	}
	return area
}