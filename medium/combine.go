package main

import (
	"fmt"
)
//NO.77
func main() {
	fmt.Println(combine(4,2))
}

//回溯算法
func combine(n int, k int) [][]int {
	var result [][]int
	var temp []int
	var dfs func(int)
	dfs = func(i int) {
		if len(temp)+(n-i+1) <k {
			return
		}
		if len(temp) == k {
			com:= make([]int, k)
			copy(com, temp)
			result = append(result, com)
			return
		}
		temp = append(temp, i)
		dfs(i+1)
		temp = temp[:len(temp)-1]
		dfs(i+1)
	}
	dfs(1)
	return result
}
