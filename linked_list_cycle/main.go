package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func hasCycle(head *ListNode) bool {
	visitedNodes := make(map[*ListNode]bool)
	currentNode := head
	for currentNode != nil {
		if visitedNodes[currentNode] {
			return true
		}
		visitedNodes[currentNode] = true
		currentNode = currentNode.Next
	}

	return false
}
