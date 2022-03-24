package main

import "fmt"

//NO.6028
func main() {
	fmt.Println(countCollisions("RLRSLL"))
}

func countCollisions(directions string) int {
	count :=0
	flag := false
	for i:=0;i<len(directions);i++ {
		if directions[i] == 'R' || directions[i] == 'S' {
			flag = true
		} else if flag {
			count++
		}
	}
	flag = false
	for i:=len(directions)-1;i>=0;i-- {
		if directions[i] == 'L' || directions[i] == 'S' {
			flag = true
		} else if flag {
			count++
		}
	}
	return count
}

