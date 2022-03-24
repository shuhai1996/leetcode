package main

import "fmt"
//NO.3
func main() {
	s:= "abcdea"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	curL :=0 //当前字符串长度
	maxL :=0 //最大字串长度
	position :=0 //字串起始下标
	m := make(map[byte]int)
	charAr := []byte(s) //字符串改成数组

	for i:=0;i<len(charAr);i++{

		if v, ok := m[charAr[i]]; ok { //判断map中是否有key，
			fmt.Println("ok")
			fmt.Println(v)
			if position <= v { //如果有并且position <= v，字串起始下标挪到重复字符后一个
				position = v+1
				curL = i - position +1
			} else {
				curL++
				if curL>maxL {
					maxL = curL
				}
			}
		} else {
			curL++
			if curL>maxL {
				maxL = curL
			}
		}
		m[charAr[i]] = i
	}
	return maxL
}

