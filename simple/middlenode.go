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
	node3.Val = 2
	node2.Next = &node3
	node4 :=ListNode{}
	node4.Val = 1
	node3.Next = &node4
	node5 :=ListNode{}
	node5.Val = 5
	node4.Next = &node5
	fmt.Println(middleNode(head).Val)
}

//快慢指针
func middleNode(head *ListNode) *ListNode {
	fast:=head
	slow:=head
	for fast!=nil && fast.Next!=nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func middleNodeV1(head *ListNode) *ListNode {
	var n []int
	var old *ListNode
	old = head
	for head!=nil {
		n = append(n, head.Val)
		head = head.Next
	}
	for i:=0;i<len(n)/2;i++ {
		old = old.Next
	}
	return old
}

type ListNode struct {
	Val int
	Next *ListNode
}

