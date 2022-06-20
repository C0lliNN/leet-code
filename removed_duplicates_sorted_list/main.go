package main

import "fmt"

func main() {
	head1 := &ListNode{1, &ListNode{1, &ListNode{2, nil}}}
	head2 := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}

	printList(deleteDuplicates(head1))
	fmt.Println("===============")
	printList(deleteDuplicates(head2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	curr := head
	advanced := curr.Next
	for advanced != nil {
		if curr.Val == advanced.Val {
			curr.Next = advanced.Next
			advanced = advanced.Next
		} else {
			curr = advanced
			advanced = advanced.Next
		}
	}

	return head
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}
