package main

import "fmt"

// NO.232
func main() {
	obj := Constructor()
	obj.Push(1)
	param2 := obj.Pop()
	param3 := obj.Peek()
	param4 := obj.Empty()
	fmt.Println(param2)
	fmt.Println(param3)
	fmt.Println(param4)
}

type MyQueue struct {
	inStack []int //进栈
	outStack []int //出栈
}


func Constructor() MyQueue {
	return MyQueue{}
}


func (q *MyQueue) Push(x int)  {
	q.inStack = append(q.inStack, x)
}


func (q *MyQueue) Pop() int {
	if len(q.outStack) ==0 {
		q.in2out()
	}
	x := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return x
}


func (q *MyQueue) Peek() int {
	if len(q.outStack) == 0 {
		q.in2out()
	}
	return q.outStack[len(q.outStack)-1]
}


func (q *MyQueue) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}

func (q *MyQueue) in2out() {
	for len(q.inStack) > 0 {
		q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
		q.inStack = q.inStack[:len(q.inStack)-1]
	}
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */