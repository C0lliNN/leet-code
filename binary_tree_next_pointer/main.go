package main

func main() {

}

 type Node struct {
	     Val int
	     Left *Node
	     Right *Node
	     Next *Node
	 }

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	connectHelper(root.Left, root.Right)
	return root
}

func connectHelper(root1, root2 *Node) {
	if root1 == nil || root2 == nil {
		return
	}

	root1.Next = root2

	connectHelper(root1.Left, root1.Right)
	connectHelper(root2.Left, root2.Right)
	connectHelper(root1.Right, root2.Left)
}