package main

import "fmt"
//NO.231
func main() {
	fmt.Println(isPowerOfTwo(0))
}

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}

	for n%2==0 {
		n = n/2
	}

	if n==1 {
		return true
	} else {
		return false
	}
}