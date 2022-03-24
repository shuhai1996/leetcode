package main

import "fmt"

//NO.206
func main() {
	head1 := ListNode{}
	head1.Val =1
	node1 := ListNode{}
	node1.Val =3
	head1.Next = &node1

	a:= reverseList(&head1)
	for a != nil {
		fmt.Println(a.Val)
		a = a.Next
	}
}
//链表反转，利用虚指针
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	curl := head //移动指针位置
	for curl!= nil {
		next :=curl.Next
		curl.Next = pre
		pre = curl
		curl = next
	}
	return pre
}

type ListNode struct {
	Val int
	Next *ListNode
}