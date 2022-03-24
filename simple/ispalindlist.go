package main

import "fmt"

//NO.234
func main() {
	node := ListNode{}
	head := &node
	head.Val = 5
	node1 :=ListNode{}
	node1.Val = 1
	head.Next = &node1
	fmt.Println(isPalindromeV3(head))
}


//反转链表
func reverse1(list *ListNode) *ListNode {
	var pre *ListNode
	curl := list //移动指针位置
	for curl!= nil {
		next :=curl.Next
		curl.Next = pre
		pre = curl
		curl = next
	}
	return pre
}

func isPalindromeV3(head *ListNode) bool {
	halfEnd:=getHalfEnd1(head)
	halfStart:=reverse1(halfEnd.Next) //	反转链表
	l:=head  //前半部分
	h:=halfStart //后半部分
	for h!=nil {
		if l.Val != h.Val {
			return false
		}
		l=l.Next
		h=h.Next
	}
	return true
}

func getHalfEnd1(list *ListNode) *ListNode {
	fast := list //快慢指针实现
	slow := list
	for fast.Next !=nil && fast.Next.Next!= nil{
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
