package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	queue := make([][]int, 0)
	return levelHelper(queue, root, 0)
}

func levelHelper(queue [][]int, root *TreeNode, depth int) [][]int {
	if root == nil {
		return queue
	}

	if depth >= len(queue) {
		queue = append(queue, make([]int, 0))
	}
	queue[depth] = append(queue[depth], root.Val)
	queue = levelHelper(queue, root.Left, depth+ 1)
	queue = levelHelper(queue, root.Right, depth+ 1)

	return queue
}