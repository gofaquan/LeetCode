package ___offer

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	r := &ListNode{}
	d := r
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			r.Next = l2
			l2 = l2.Next
		} else {
			r.Next = l1
			l1 = l1.Next
		}

		r = r.Next
	}

	for l1 != nil {
		r.Next = l1
		r = r.Next
		l1 = l1.Next
	}

	for l2 != nil {
		r.Next = l2
		l2 = l2.Next
		r = r.Next

	}

	return d.Next
}
