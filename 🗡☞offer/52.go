package ___offer

// 朴素
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	m := map[*ListNode]bool{}

	for headA != nil && headB != nil {
		if m[headA] == true {
			return headA
		}
		m[headA] = true
		headA = headA.Next

		if m[headB] == true {
			return headB
		}
		m[headB] = true
		headB = headB.Next
	}

	for headA != nil {
		if m[headA] == true {
			return headA
		}
		m[headA] = true

		headA = headA.Next
	}

	for headB != nil {
		if m[headB] == true {
			return headB
		}
		m[headB] = true

		headB = headB.Next
	}

	return nil
}

// 优解
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}

		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
		// pA ,pB 互换一次后会同时走到最后一个节点
		// 2 -> 3
		// 1

		// 对于上方这个例子，不相交会同时走向 nil 就是 相等了 a = b = nil
	}

	return pA
}
