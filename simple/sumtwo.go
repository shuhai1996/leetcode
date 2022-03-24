package main

func twoSum(nums []int, target int) []int {
	var m =0
	var result []int
	for m<len(nums) {
		var n =0
		for n<len(nums){
			if m != n && nums[m]+ nums[n] == target {
				result = append(result, m)
				result = append(result, n)
				return result
			}
			n++
		}
		m++
	}
	return result
}