package main

import (
	"fmt"
)

func main() {
	fmt.Println(myAtoi("00000-42a1234"))
}

func myAtoi(s string) int {
	strArr:= []byte(s)
	var newStrArr []int
	flag:=1 // 默认为正数
	symbol :=0 //有符号了
	for i:=0;i<len(strArr);i++{
		if len(newStrArr) ==0{
			if symbol>0 {
				if int(strArr[i]) -48 <0 || int(strArr[i]) -48>9 {
					break
				}
			}
			if strArr[i] ==' '{
				continue
			} else if strArr[i] == '-' {
				flag =0 //为负数
				symbol = 1
				continue
			} else if strArr[i] == '+' {
				flag =1 //为负数
				symbol = 1
				continue
			} else if strArr[i] == '0' {
				symbol = 1
				continue
			}
		}

		if int(strArr[i]) -48 <0 || int(strArr[i]) -48>9 {
			break
		}
		newStrArr = append(newStrArr, int(strArr[i])-48)
	}
	var z int
	for i:=0;i<len(newStrArr);i++ {
		if z >= 214748364{ //超过限制
			if z== 214748364 && newStrArr[i]<=7 {
				z= z*10+newStrArr[i]
			} else if flag == 0 {
				z= 2147483648
			} else {
				z= 2147483647
			}
		} else {
			z= z*10+newStrArr[i]
		}
	}
	if flag == 0 {
		z = -z
	}
	return z
}

