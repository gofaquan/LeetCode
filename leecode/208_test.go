package leecode

type Trie struct {
	child [26]*Trie
	ended bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, value := range word {
		value -= 'a'
		if node.child[value] == nil {
			node.child[value] = &Trie{}
		}

		node = node.child[value]
	}

	node.ended = true
}
func (this *Trie) SearchPrefix(prefix string) *Trie {
	node := this
	for _, value := range prefix {
		value -= 'a'
		if node.child[value] == nil {
			return nil
		}
		node = node.child[value]
	}
	return node
}

func (this *Trie) Search(word string) bool {
	node := this.SearchPrefix(word)
	return node != nil && node.ended
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.SearchPrefix(prefix)
	return node != nil
}
