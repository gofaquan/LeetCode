package main

import "fmt"

var (
	head = -1 // -1 代表空节点
	idx  = 0  // 下一个有效数的下标，初始值默认是第一个元素，即下标 0
)

var element = [10]int{}
var next = [10]int{}

// AddToHead 插入到头节点
func AddToHead(x int) {
	// 给新空间 先赋值
	element[idx] = x // 起始下标赋值
	// 赋值后 确定插入位置
	// 进行 插入
	next[idx] = head // 新空间的 next 数组 指向 前一个节点，第一个插入的则是空节点
	head = idx       // 插入完新的前节点就是当前被插入的节点
	// 后移插入下标 确定下个元素进来的位置
	idx++
}

// Add 指定 插入下标 插入
func Add(k, x int) {
	// 给新空间 先赋值
	element[idx] = x // 起始下标赋值
	// 赋值后 确定插入位置
	// 进行 插入
	next[idx] = next[k] // 新节点的 next 是 位置 k 节点 的 next
	next[k] = idx       // k 节点 next 是 idx
	// 后移插入下标 确定下个元素进来的位置
	idx++
}

// Remove 删除节点
func Remove(k int) {
	next[k] = next[next[k]] // k 的 next 指向 k 的 next 的 next ，跳过 k
}

func printList() {
	for i := head; i != -1; i = next[i] {
		fmt.Printf("%d ", element[i])
	}
}

func main() {
	AddToHead(1)
	AddToHead(2)
	AddToHead(3)
	AddToHead(4)
	fmt.Println(element)
	Add(1, 20)
	fmt.Println(element)
	//Remove(1)
	//fmt.Println(element)

	fmt.Println(next)
	// 双向链表
	//list.New()
	printList()
}
