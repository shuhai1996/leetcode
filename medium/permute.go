package main

import "fmt"

//NO.46
func main() {
	fmt.Println(permute([]int{1,2,3}))
}
//回溯算法
func permute(nums []int) [][]int {
	var result [][]int
	var temp []int
	var dfs func()
	//将用过节点进行标记的哈希表
	visited :=make(map[int]bool)
	l := len(nums)
	dfs = func() {
		if len(temp) == l { //终止条件
			com:= make([]int, l)
			copy(com, temp)
			result = append(result, com)
			return
		}
		for i:=0;i<l;i++ {
			if visited[nums[i]] {
				continue
			}
			visited[nums[i]] = true
			temp = append(temp, nums[i])
			dfs()//递归
			visited[nums[i]] = false //回溯
			temp = temp[:len(temp)-1]
		}
	}
	dfs()
	return result
}