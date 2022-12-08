package leecode

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 不进位情况下就是错的，多了值为0的 尾节点
func TestAddTwo(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	fmt.Println(addTwoNumbers(l1, l2).Next.Next.Val)
	//fmt.Println(addTwoNumbers(l1, l2).Next.Next.Next.Val)
}

func addTwoNumbers2(l1, l2 *ListNode) *ListNode {
	tail := &ListNode{}
	head := tail
	carry := 0
	first := true

	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10

		if first {
			tail.Val = sum
			first = false
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}

	}

	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}

	return head
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	tail := head
	t := 0
	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		t += v1 + v2
		if head == nil {
			head = &ListNode{Val: t % 10}
			tail = head
		} else {
			tail.Next = &ListNode{Val: t % 10}
			tail = tail.Next
		}

		t /= 10
	}

	if t > 0 {
		tail.Next = &ListNode{Val: t}
	}

	return
}
