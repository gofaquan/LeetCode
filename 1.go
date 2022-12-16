package main

type ListNode struct {
	Val  int
	Pre  *ListNode
	Next *ListNode
}

type MyLinkedList struct {
	Size int
	Head *ListNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{0, &ListNode{}}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Size {
		return -1
	}

	cur := this.Head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.Size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Size {
		return
	}

	this.Size++
	pre := this.Head
	index = max(index, 0)
	for i := 0; i < index; i++ {
		pre = pre.Next
	}
	toAdd := &ListNode{
		Val:  val,
		Pre:  pre,
		Next: pre.Next,
	}
	pre.Next = toAdd
	toAdd.Next.Pre = toAdd
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > this.Size {
		return
	}
	this.Size--
	pre := this.Head

	for i := 0; i < index; i++ {
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
	pre.Next.Next.Pre = pre
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
