package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}
	node := removeElements(head, 2)
	printList(node)
}
func removeElements(head *ListNode, val int) *ListNode {
	newHead := head
	current := newHead
	var prev *ListNode
	for current != nil {
		if current.Val == val {
			if prev == nil {
				newHead = newHead.Next
				current = newHead
			} else {
				prev.Next = current.Next
				current = current.Next
			}
		} else {
			prev = current
			current = current.Next
		}

	}

	return newHead
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}