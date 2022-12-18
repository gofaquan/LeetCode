package ___offer

// 保存差值栈

type MinStack struct {
	min   int
	stack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		min:   0,
		stack: []int{},
	}
}

func (this *MinStack) Push(x int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, 0)
		this.min = x
	} else {
		this.stack = append(this.stack, x-this.min)
		if x < this.min {
			this.min = x
		}
	}

}

func (this *MinStack) Pop() {
	if this.stack[len(this.stack)-1] < 0 {
		this.min = this.min - this.stack[len(this.stack)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	if this.stack[len(this.stack)-1] > 0 {
		return this.stack[len(this.stack)-1] + this.min
	}

	return this.min
}

func (this *MinStack) Min() int {
	return this.min
}
