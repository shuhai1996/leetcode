package main

import (
	"fmt"
)

func main() {
	fmt.Println(strStr("a","a"))
}

func strStr(haystack string, needle string) int {
	if needle =="" {
		return 0
	}
	length := len(needle)
	outer:
		for i:=0;i+length<=len(haystack);i++ {
			for j:=range needle {
				if haystack[i+j] != needle[j] {
					continue outer
				}
			}
			return i
		}
		return -1
}


