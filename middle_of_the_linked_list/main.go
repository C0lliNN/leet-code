package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	head2 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}

	fmt.Println(middleNode2(head1).Val)
	fmt.Println(middleNode2(head2).Val)
}

func middleNode(head *ListNode) *ListNode {
	nodes := make([]*ListNode, 0)

	current := head
	for current != nil {
		nodes = append(nodes, current)
		current = current.Next
	}

	return nodes[len(nodes) / 2]
}

func middleNode2(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
