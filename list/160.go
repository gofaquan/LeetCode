package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// hashmap
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	m := map[*ListNode]bool{}
	for headA != nil {
		m[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if m[headB] == true {
			return headB
		}
		headB = headB.Next
	}

	return nil
}

// 双指针
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
