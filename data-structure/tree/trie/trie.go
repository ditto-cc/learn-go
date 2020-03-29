package trie

type Trie struct {
	root *Node
	size int
}

func NewTrie() *Trie {
	return &Trie{
		root: NewNode(),
	}
}

type Node struct {
	next   map[rune]*Node
	isWord bool
}

func NewNode() *Node {
	return &Node{}
}
