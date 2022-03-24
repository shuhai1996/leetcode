package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var a int
	var b int
	var head = new(ListNode)
	head.Val = 0
	var tail *ListNode // 尾插法
	var re []int
	for {
		a = 0
		b = 0
		if l1 != nil {
			a = l1.Val
			l1=l1.Next  //移动指针
		}
		if l2 != nil {
			b = l2.Val
			l2=l2.Next
		}

		if a+b >=10 {
			if l1!= nil {
				l1.Val = l1.Val+1
			} else {
				var n = new(ListNode)
				n.Val =1
				l1 = n
			}
		}

		re = append(re, (a+b)%10)
		if l1 == nil && l2 ==nil {
			break
		}
	}

	for i := len(re); i > 0; i-- {
		var node = ListNode{
			Val: re[i-1],
		}
		node.Next = tail
		tail = &node
	}
	return tail

}
