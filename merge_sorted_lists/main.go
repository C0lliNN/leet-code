package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	curr1, curr2 := l1, l2
	var head, curr3 *ListNode
	for curr1 != nil && curr2 != nil {
		var selectedNode *ListNode
		if curr1.Val <= curr2.Val {
			selectedNode = curr1
			curr1 = curr1.Next
		} else {
			selectedNode = curr2
			curr2 = curr2.Next
		}
		if curr3 == nil {
			curr3 = selectedNode
			head = curr3
		} else {
			curr3.Next = selectedNode
			curr3 = selectedNode
		}
	}

	if curr1 != nil {
		curr3.Next = curr1
	}
	if curr2 != nil {
		curr3.Next = curr2
	}

	return head
}
