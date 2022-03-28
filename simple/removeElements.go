package main

import "fmt"

func main() {
	node := ListNode{}
	head := &node
	head.Val = 5
	node1 :=ListNode{}
	node1.Val = 1
	head.Next = &node1
	node2 :=ListNode{}
	node2.Val =8
	node1.Next = &node2
	node3 :=ListNode{}
	node3.Val = 8
	node2.Next = &node3
	node4 :=ListNode{}
	node4.Val = 1
	node3.Next = &node4
	node5 :=ListNode{}
	node5.Val = 5
	node4.Next = &node5
	fmt.Println(removeElements(head,4))
}

func removeElements(head *ListNode, val int) *ListNode {
	curr:= &ListNode{Next: head}

	for temp:=curr; temp.Next != nil;{
		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}
	return curr.Next
}

type ListNode struct {
	Val int
	Next *ListNode
}
