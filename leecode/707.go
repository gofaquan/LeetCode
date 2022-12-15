package leecode

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	Size int
	Head *ListNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{0, &ListNode{}}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.Size {
		return -1
	}
	cur := l.Head
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (l *MyLinkedList) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *MyLinkedList) AddAtTail(val int) {
	l.AddAtIndex(l.Size, val)
}

func (l *MyLinkedList) AddAtIndex(index, val int) {
	if index > l.Size {
		return
	}
	index = max(index, 0)
	l.Size++
	pred := l.Head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	toAdd := &ListNode{val, pred.Next}
	pred.Next = toAdd
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.Size {
		return
	}
	l.Size--
	pre := l.Head
	for i := 0; i < index; i++ {
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
