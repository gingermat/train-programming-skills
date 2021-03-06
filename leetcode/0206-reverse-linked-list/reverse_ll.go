package reverse_linked_list

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	reversed := reverseList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return reversed
}
