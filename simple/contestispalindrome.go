package main

import (
	"fmt"
)

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
	fmt.Println(isPalindV1(head))

	//fmt.Println(isPalind([]int{1,2,2,3,1}, 2))
}

type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var nums []int
	for head != nil{
		nums = append(nums, head.Val)
		head = head.Next
	}
	fmt.Println(nums)
	for position:=0;position< len(nums); position++ {
		if isPalind(nums, position) {
			fmt.Println(position)
			return true
		}
	}
	return false
}

func isPalind(nums []int, position int) bool {
	i:=0
	l := 0
	r := 0
	for i<=(len(nums)-1)/2 {
		l = i
		if position >= (len(nums)-1-l){
			r = len(nums)-2-i
		} else {
			r = len(nums)-1-i
		}
		if nums[l] != nums[r] {
			return false
		}
		i++
	}
	return true
}

//反转链表
func reverse(list *ListNode) *ListNode {
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

func isPalindV1(head *ListNode) bool {
	halfEnd:=getHalfEnd(head)
	halfStart:=reverse(halfEnd.Next) //	反转链表
	for halfStart!=nil {
		fmt.Println(halfStart.Val)
		halfStart = halfStart.Next
	}
	l:=head  //前半部分
	h:=halfStart //后半部分
	for h!=nil {
		if l.Val != h.Val {
			return false
		}
	}
	return true
}

func getHalfEnd(list *ListNode) *ListNode {
	fast := list //快慢指针实现
	slow := list
	for fast.Next !=nil && fast.Next.Next!= nil{
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
