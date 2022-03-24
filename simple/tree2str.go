package main

import (
	"fmt"
	"strconv"
)

//NO.606
func main() {

}

func tree2str(root *TreeNode) string {
	if root==nil {
		return ""
	}
	if root.Left == nil && root.Right == nil {
		return strconv.Itoa(root.Val)
	} else if root.Right == nil {
		return fmt.Sprintf("%d(%s)", root.Val, tree2str(root.Left))
	} else {
		return fmt.Sprintf("%d(%s)(%s)", root.Val, tree2str(root.Left), tree2str(root.Right))
	}
}


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

