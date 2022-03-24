package main

import "fmt"

func main() {
	fmt.Println(reverse(7453847412))
}

func reverse(x int) int {
	var y []int
	if x==0 {
		return 0
	}
	temp := x
	if x<0 {
		x = -x
	}
	if x/10 <=0 {
		if temp<0 {
			return -x
		}
		return x
	}
	i :=0
	flag := 0
	for x >0 {
		if x%10 ==0 {
			if flag!=i {
				y = append(y, x%10)
			}
			flag++
		} else {
			y = append(y, x%10)
		}
		x = x/10
		i++
	}

	if len(y) ==10 && y[0]>=5 {
		return 0
	}
	if len(y)==1 {
		if temp<0 {
			return -y[0]
		}
		return y[0]
	}
	fmt.Println(y)
	var z int
	for i:=0;i<len(y);i++{
		if z >= 214748364{ //超过限制
			if z== 214748364 && y[i]<=7 {
				z= z*10+y[i]
			} else {
				return 0
			}
		} else {
			z= z*10+y[i]
		}

	}
	if temp<0 {
		return -z
	}
	return z
}