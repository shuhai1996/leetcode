package main

import "fmt"
//NO.21
func main() {
	head1 := ListNode{}
	head1.Val =1
	node1 := ListNode{}
	node1.Val =3
	head1.Next = &node1

	head2 := ListNode{}
	head2.Val=2
	node2 := ListNode{}
	node2.Val =4
	head2.Next = &node2
	a:= mergeTwoLists(&head1, &head2)
	for a != nil {
		fmt.Println(a.Val)
		a = a.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1==nil {
		return list2
	}else if list2==nil {
		return list1
	}else if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}



	//归并
func mergeTwoListsV1(list1 *ListNode, list2 *ListNode) *ListNode {
	var node = ListNode{}
	head := &node
	for  {
		if list1 !=nil && list2 !=nil {
			if list1.Val <= list2.Val {
				head.Next = list1
				list1 = list1.Next
			} else {
				head.Next = list2

				list2 = list2.Next
			}
			head = head.Next
		}
		if list1 == nil {
			head.Next = list2
			break
		}
		if list2 == nil {
			head.Next = list1
			break
		}
	}
	return node.Next //返回起始位置
}

type ListNode struct {
	Val int
	Next *ListNode
}