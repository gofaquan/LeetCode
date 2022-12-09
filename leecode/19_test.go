package leecode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var length = 1
	test := head
	for test != nil {
		test = test.Next
		length++
	}

	dummy := &ListNode{0, head}
	test = dummy

	for i := 0; i < length-n-1; i++ {
		test = test.Next
	}
	test.Next = test.Next.Next

	return dummy.Next
}
