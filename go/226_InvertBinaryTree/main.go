package main

import "fmt"

func main() {
	var root *TreeNode
	temp := invertTree(root)
	fmt.Println(temp)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	invertNode(root)
	return root
}

func invertNode(node *TreeNode) {
	if node != nil {
		node.Left, node.Right = node.Right, node.Left
		invertNode(node.Left)
		invertNode(node.Right)
	}
}

func invertTreeTop(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = root.Right, root.Left
		invertTree(root.Left)
		invertTree(root.Right)
	}
	return root
}
