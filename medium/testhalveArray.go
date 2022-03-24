package main

import (
	"container/heap"
	"fmt"
)


//NO.6022
func main() {
	fmt.Println(halveArray([]int{5,19,8,1}))
}

type IntHeap []float64

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(float64))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}



//优先队列解法
func halveArray(nums []int) int {
	var s float64
	h:= &IntHeap{}
	heap.Init(h)
	for _,x:=range nums{
		s += float64(x)
		heap.Push(h, float64(x))
	}

	var now float64
	ans:=0
	for now*2 <s {
		pop := heap.Pop(h).(float64)
		now += pop/2
		heap.Push(h, pop/2)
		ans++
	}
	return ans
}