package leecode

func _234isPalindrome(head *ListNode) bool {
	var stack []int

	for ; head != nil; head = head.Next {
		stack = append(stack, head.Val)
	}

	for i := 0; i < len(stack)/2+1; i++ {
		if stack[i] != stack[len(stack)-1-i] {
			return false
		}
	}
	return true
}

func isPalindrome1(head *ListNode) bool {
	frontPointer := head
	var recursivelyCheck func(*ListNode) bool
	recursivelyCheck = func(curNode *ListNode) bool {
		if curNode != nil {
			if !recursivelyCheck(curNode.Next) {
				return false
			}
			if curNode.Val != frontPointer.Val {
				return false
			}
			frontPointer = frontPointer.Next
		}

		return true
	}
	return recursivelyCheck(head)
}
