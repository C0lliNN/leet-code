package main

 type ListNode struct {
	     Val int
	     Next *ListNode
	 }

func main() {

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	// the 2nd pointer will advance only after
	// we processed the first n nodes
	first, second := dummy, dummy

	var count int
	for first.Next != nil {
		first = first.Next

		count++
		if count > n {
			second = second.Next
		}
	}

	second.Next = second.Next.Next

	return dummy.Next
}
