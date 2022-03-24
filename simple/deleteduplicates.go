package main

import "fmt"

//NO.83
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
	node3.Val = 3
	node2.Next = &node3
	node4 := ListNode{}
	node4.Val = 4
	node3.Next = &node4
	pre :=deleteDuplicates(head)
	for pre != nil {
		fmt.Println(pre.Val)
		pre = pre.Next
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	pre := head
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	for head.Next != nil{
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return pre
}


type ListNode struct {
	Val int
	Next *ListNode
}