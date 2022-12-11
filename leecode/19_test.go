package leecode

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	var length = 0
	test := head
	for test != nil {
		test = test.Next
		length++
	}

	dummy := &ListNode{0, head}
	test = dummy

	for i := 0; i < length-n; i++ {
		test = test.Next
	}
	test.Next = test.Next.Next

	return dummy.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p, q := head, head
	if head == nil || head.Next == nil {
		return nil
	}

	j := 1
	for q.Next != nil {
		q = q.Next
		j++
	}

	if j-n == 0 {
		return head.Next
	}

	for i := 0; ; i++ {
		if i == j-n-1 {
			p.Next = p.Next.Next
			break
		}
		p = p.Next
	}

	return head
}
