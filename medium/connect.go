package main

func main() {

}

func connect(root *Node) *Node {
	if root ==nil {
		return root
	}
	queue:= []*Node{root}//初始化队列同时将第一层节点加入队列中，即根节点
	for len(queue)>0 {
		tmp:= queue
		queue =nil
		for index, node:=range tmp{
			if index+1  < len(tmp) { // 连接起来
				node.Next = tmp[index+1]
			}

			//加入下一层节点
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}