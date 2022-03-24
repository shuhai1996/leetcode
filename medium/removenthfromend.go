package main

import (
	"fmt"
)

//NO.19
func main() {
	node := ListNode{}
	head := &node
	head.Val = 1
	node1 := ListNode{}
	node1.Val = 2
	head.Next = &node1
	node2 := ListNode{}
	node2.Val = 3
	node1.Next = &node2
	node3 := ListNode{}
	node3.Val = 4
	node2.Next = &node3
	node4 := ListNode{}
	node4.Val = 5
	node3.Next = &node4
	//node5 := ListNode{}
	//node5.Val = 99
	//node4.Next = &node5
	k:= removeNthFromEnd(head, 3)
	for k!=nil {
		fmt.Println(k.Val)
		k = k.Next
	}
	fmt.Println()
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmp:= &ListNode{0,head}
	fast,slow := head, tmp

	for i:=0;i<n;i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast= fast.Next
	}
	slow.Next = slow.Next.Next
	return tmp.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
