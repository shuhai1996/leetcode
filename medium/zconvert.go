package main

import (
	"fmt"
)

func main() {
	fmt.Println(convert("ABCDEEFG", 2))
}


func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	strArr := []byte(s)
	currentLineI := 0 // 当前列-1
	res:=make([][]byte,numRows) //结果字符串map
	flag := true // 标记是往下还是往上排列，true的时候往下，反之往上
	for i:=0;i<len(strArr);i++ {
		if currentLineI >= len(res) {
			res[currentLineI] = make([]byte,1)
		}
		res[currentLineI] = append(res[currentLineI], strArr[i])
		if flag {
			currentLineI ++
			if currentLineI >= numRows-1 {//往下还是往上排列
				flag = changeFlag(flag)
			}
		}else {
			currentLineI --
			if currentLineI ==0 {
				flag = changeFlag(flag)
			}
		}
	}
	var result []byte
	for j:=0; j<len(res);j++ {
		result = append(result, res[j]...)
	}
	return string(result)
}

func changeFlag(flag bool) bool {
	if flag {
		return false
	}
	return true
}