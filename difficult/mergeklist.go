package main

func main() {

}

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var head *ListNode
	if len(lists) == 0 {
		return head
	}

	head = lists[0]
	if len(lists) > 1 {
		for i:=1;i<len(lists);i++ {
			head = mergeTwoLists(head, lists[i])
		}
	}
	return head
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