package trie

func (t *Trie) Add(word string) {
	p := t.root
	for _, c := range word {
		if _, ok := p.next[c]; !ok {
			p.next[c] = NewNode()
		}
		p = p.next[c]
	}
	p.isWord = true
	t.size++
}

func (t *Trie) Search(word string) bool {
	p := t.root
	for _, c := range word {
		if _, ok := p.next[c]; !ok {
			return false
		}
		p = p.next[c]
	}
	return p.isWord
}

func (t *Trie) IsPrefix(prefix string) bool {
	p := t.root
	for _, c := range prefix {
		if _, ok := p.next[c]; !ok {
			return false
		}
		p = p.next[c]
	}
	return true
}

func (t *Trie) Match(pattern string) bool {
	return t.root.match(pattern, 0)
}

func (node *Node) match(pattern string, i int) bool {
	if i == len(pattern) {
		return node.isWord
	}

	if pattern[i] == '.' {
		for _, v := range node.next {
			if v.match(pattern, i+1) {
				return true
			}
		}
	} else {
		if next, ok := node.next[rune(pattern[i])]; ok {
			return next.match(pattern, i+1)
		}
	}
	return false
}

func (t *Trie) Size() int {
	return t.size
}
