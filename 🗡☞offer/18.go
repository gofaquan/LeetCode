package ___offer

func deleteNode(head *ListNode, val int) *ListNode {
	cur := head
	pre := head

	if head.Val == val {
		return head.Next
	}

	for cur != nil && cur.Val != val {
		cur = cur.Next
	}

	for pre.Next != cur {
		pre = pre.Next
	}

	pre.Next = cur.Next

	return head
}
