package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, math.MinInt, math.MaxInt)
}

func isValidBSTHelper(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val >= max || root.Val <= min {
		return false
	}

	return isValidBSTHelper(root.Left, min, root.Val) && isValidBSTHelper(root.Right, root.Val, max)
}