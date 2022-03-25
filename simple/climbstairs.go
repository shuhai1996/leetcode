package main
//NO.70
import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(climbStairs(3))
}
//f(n) = f(n - 1) + f(n - 2)f(n)=f(n−1)+f(n−2)
//直接通项公式求
func climbStairs(n int) int {
	sqrt5 := math.Sqrt(5)
	pow1 := math.Pow((1+sqrt5)/2, float64(n+1))
	pow2 := math.Pow((1-sqrt5)/2, float64(n+1))
	return int(math.Round((pow1 - pow2) / sqrt5))
}
