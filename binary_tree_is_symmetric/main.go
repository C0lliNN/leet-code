package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
	}

	fmt.Println(isSymmetric(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return root == nil  || compareSiblings(root.Left, root.Right)
}

func compareSiblings(node1, node2 *TreeNode) bool {
	if node1 == nil || node2 == nil {
		return node1 == node2
	}

	if node1.Val != node2.Val {
		return false
	}

	return compareSiblings(node1.Left, node2.Right) && compareSiblings(node1.Right, node2.Left)
}