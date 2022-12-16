package main

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func isPalindrome(head *ListNode) bool {

	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		//fast.Next = fast.Next.Next
		fast = fast.Next.Next
		slow = slow.Next
	}

	cur := slow.Next
	var pre *ListNode
	for cur != nil {
		p := cur.Next
		cur.Next = pre
		pre = cur
		cur = p
	}

	tmp := head
	for tmp != nil && pre != nil {
		if tmp.Val != pre.Val {
			return false
		}
		tmp = tmp.Next
		pre = pre.Next
	}

	return true
}

func TestIsPalindrome(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  0,
				Next: nil,
			},
		},
	}

	isPalindrome(head)
}
