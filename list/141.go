package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// hashmap
func hasCycle1(head *ListNode) bool {
	m := map[*ListNode]bool{}
	for head != nil {
		if m[head] == true {
			return true
		}

		m[head] = true
		head = head.Next
	}

	return false
}

// 快慢指针
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
