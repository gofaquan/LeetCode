package ___offer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	m := map[*Node]*Node{}
	cur := head
	for cur != nil {
		m[cur] = &Node{
			Val: cur.Val,
		}

		cur = cur.Next
	}

	cur = head
	for cur != nil {
		m[cur].Next = m[cur.Next]
		m[cur].Random = m[cur.Random]
		cur = cur.Next
	}

	return m[head]
}
