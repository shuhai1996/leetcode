package main

func main() {

}
//NO.100
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q==nil {
		return true
	}
	if p == nil || q== nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}