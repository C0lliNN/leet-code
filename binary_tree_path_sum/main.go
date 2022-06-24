package main

func main() {

}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	sum := targetSum - root.Val
	if sum == 0 && isLeaf(root){
		return true
	}

	if hasPathSum(root.Left, sum) {
		return true
	}

	return hasPathSum(root.Right, sum)
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
