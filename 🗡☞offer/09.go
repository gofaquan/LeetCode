package ___offer

type CQueue struct {
	inStack, outStack []int
}

func Constructor() CQueue {
	return CQueue{
		inStack:  []int{},
		outStack: []int{},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.inStack = append(this.inStack, value)
}

func (this *CQueue) DeleteHead() int {
	return this.convert()
}

func (this *CQueue) convert() int {
	if len(this.inStack) == 0 {
		return -1
	}

	this.outStack = this.inStack[1:]
	val := this.inStack[0]
	this.inStack = this.outStack
	return val
}
