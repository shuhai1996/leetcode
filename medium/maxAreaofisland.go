package main

import "fmt"

//NO.695
func main() {
	fmt.Println(maxAreaOfIsland([][]int{{0,0,1,0,0,0,0,1,0,0,0,0,0},{0,0,0,0,0,0,0,1,1,1,0,0,0},{0,1,1,0,1,0,0,0,0,0,0,0,0},{0,1,0,0,1,1,0,0,1,0,1,0,0},{0,1,0,0,1,1,0,0,1,1,1,0,0},{0,0,0,0,0,0,0,0,0,0,1,0,0},{0,0,0,0,0,0,0,1,1,1,0,0,0},{0,0,0,0,0,0,0,1,1,0,0,0,0}}))
}

func maxAreaOfIsland(grid [][]int) int {
	max:=0
	for i:=0;i<len(grid);i++ {
		for j:=0;j<len(grid[0]);j++{
			if grid[i][j] == desColor {
				area := getArea(grid, i, j)
				if area > max {
					max =area
				}
			}
		}
	}
	return max
}

func getArea(grid [][]int, posX int, posY int) int{
	var queue [][]int
	queue = append(queue, []int{posX, posY})
	grid[posX][posY] = newColor
	boardX,boardY := len(grid), len(grid[0])
	for i:=0;i<len(queue);i++ {
		for j:=0;j<4;j++ {
			posX = queue[i][0] + dx[j]
			posY = queue[i][1] + dy[j]
			if posX >=0 && posX < boardX && posY >=0 && posY < boardY && grid[posX][posY] == desColor {
				queue =append(queue, []int{posX, posY})
				grid[posX][posY] = newColor
			}
		}
	}

	return  len(queue)
}

var ( //定义向下、右、上、左移动的四种情况 ，x增减是改变上下、y增减是改变左右
	dx = []int{1, 0, 0, -1}
	dy = []int{0, 1, -1, 0}
	desColor =1
	newColor =2
)