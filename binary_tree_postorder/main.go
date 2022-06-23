package main

func main() {

}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}

	result := make([]int, 0)
	if root.Left != nil {
		result = postorderTraversal(root.Left)
	}

	if root.Right != nil {
		result = append(result, postorderTraversal(root.Right)...)
	}

	return append(result, root.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
