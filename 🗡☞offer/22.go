package ___offer

func getKthFromEnd(head *ListNode, k int) *ListNode {
	l := head
	var length int
	for l != nil {
		l = l.Next
		length++
	}

	for i := 0; i < length-k; i++ {
		head = head.Next
	}

	return head
}

func getKthFromEnd2(head *ListNode, k int) *ListNode {
	f := head
	b := head
	for i := 0; i < k-1; i++ {
		f = f.Next
	}

	for f.Next != nil {
		f = f.Next
		b = b.Next
	}
	return b
}
