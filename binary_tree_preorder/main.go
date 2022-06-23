package main

import "fmt"

func main() {
	root := &TreeNode{Val: 3, Left: &TreeNode{1, nil, nil}, Right: &TreeNode{2, nil, nil}}
	fmt.Println(preorderTraversal(root))
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}

	result := make([]int, 0)
	result = append(result, root.Val)

	if root.Left != nil {
		result = append(result, preorderTraversal(root.Left)...)
	}

	if root.Right != nil {
		result = append(result, preorderTraversal(root.Right)...)
	}

	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
