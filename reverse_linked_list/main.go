package main

import "fmt"

func main() {
	head1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	head2 := &ListNode{1, &ListNode{2, nil}}

	printList(reverseList1(head1))
	fmt.Println("=============")
	printList(reverseList1(head2))
	fmt.Println("=============")
	printList(reverseList2(head1))
	fmt.Println("=============")
	printList(reverseList2(head2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	values := make([]int, 0)

	curr := head
	for curr != nil {
		values = append(values, curr.Val)
		curr = curr.Next
	}

	newList := &ListNode{values[len(values)-1], nil}
	curr = newList
	for i := len(values) - 2; i >= 0; i-- {
		curr.Next = &ListNode{values[i], nil}
		curr = curr.Next
	}

	return newList
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	var newHead *ListNode
	for head != nil {
		next := head.Next
		head.Next = newHead
		newHead = head
		head = next
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
