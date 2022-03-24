package main

func main() {

}

//NO.94
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) { //闭包
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return res
}



type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}