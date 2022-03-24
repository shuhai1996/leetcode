package main

import "fmt"
//NO.278
func main() {

	fmt.Println(firstBadVersion(6))
}

func firstBadVersion(n int) int {
	l:=1
	h:=n
	for l<h {
		half := (h+l)/2
		if isBadVersion(half) {
			h = half //收缩边界
		}else {
			l = half+1
		}
	}
	return l //	求边界
}

func isBadVersion(a int) bool {
	if a>= 4 {
		return true
	} else {
		return false
	}
}