package main
//NO.167
import "fmt"

func main() {
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
}

//双指针解法
func twoSum(numbers []int, target int) []int {
	result := make([]int, 2)
	board:= len(numbers)-1
	l:=0 // 左指针
	r:=board // 右指针始终要大于左指针
	for l<r {
		if numbers[l] + numbers[r] == target {
			result[0] = l+1
			result[1] = r+1
			break
		}else if numbers[l] + numbers[r] > target {
			r--
		}else {
			l++
		}
	}
	return result
}