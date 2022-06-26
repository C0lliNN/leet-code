package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func findTarget(root *TreeNode, k int) bool {
	values := make(map[int]bool)

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack) - 1]
		values[node.Val] = true

		stack = stack[:len(stack) - 1]

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}


	for value := range values {
		// Do not include the same value twice
		if k / 2 != value && values[k - value] {
			return true
		}
	}

	return false
}
