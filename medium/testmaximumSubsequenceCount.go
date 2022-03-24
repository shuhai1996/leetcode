package main

import (
	"fmt"
	"math"
)

//NO.6021
func main() {
	fmt.Println(maximumSubsequenceCount("mpmp",
	"mp"))
}

func maximumSubsequenceCount(text string, pattern string) int64 {
	var res int64
	var x,y int64
	for i:=0;i<len(text);i++ {
		if text[i] == pattern[1] {
			res+=x
		}
		if text[i] == pattern[0] {
			x++
		}
		if text[i] == pattern[1] {
			y++
		}
	}
	return res + int64(math.Max(float64(x), float64(y)))
}