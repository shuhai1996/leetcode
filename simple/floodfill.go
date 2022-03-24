package main

import "fmt"

//NO.733
func main() {
	fmt.Println(floodFill([][]int{{1,1,1},{1,1,0},{1,0,1}},1,1,2))
}

//队列实现广度优先搜索
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	desColor := image[sr][sc]
	if desColor == newColor {
		return image
	}
	boardX := len(image)
	boardY := len(image[0])
	posX:=0
	posY:=0
	var queue [][]int
	queue = append(queue, []int{sr, sc})
	image[sr][sc] = newColor
	for i:=0;i<len(queue); i++ {
		for j:=0;j<4;j++ {
			posX = queue[i][0]+ dx[j]
			posY = queue[i][1]+ dy[j]
			if posX>=0 && posX< boardX && posY>=0 && posY< boardY && image[posX][posY] == desColor{
				image[posX][posY] = newColor
				queue = append(queue, []int{posX, posY})
			}
		}
	}
	return image
}

var ( //定义向下、右、上、左移动的四种情况 ，x增减是改变上下、y增减是改变左右
	dx = []int{1, 0, 0, -1}
	dy = []int{0, 1, -1, 0}
)


