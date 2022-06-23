package main

func main() {

}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}

	var result []int
	if root.Left != nil {
		result = inorderTraversal(root.Left)
	}

	result = append(result, root.Val)

	if root.Right != nil {
		result = append(result, inorderTraversal(root.Right)...)
	}

	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
