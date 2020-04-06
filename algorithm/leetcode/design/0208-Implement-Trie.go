package design

/**
Implement a trie with insert, search, and startsWith methods.

Example:

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // returns true
trie.search("app");     // returns false
trie.startsWith("app"); // returns true
trie.insert("app");
trie.search("app");     // returns true
Note:

You may assume that all inputs are consist of lowercase letters a-z.
All inputs are guaranteed to be non-empty strings.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-trie-prefix-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type trieNode struct {
	next   map[byte]*trieNode
	isWord bool
}

func newTrieNode() *trieNode {
	return &trieNode{isWord: false, next: map[byte]*trieNode{}}
}

type Trie struct {
	root *trieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{root: newTrieNode()}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	p := this.root
	for i, n := 0, len(word); i < n; i++ {
		if next, ok := p.next[word[i]]; !ok {
			p.next[word[i]] = newTrieNode()
			p = p.next[word[i]]
		} else {
			p = next
		}
	}
	p.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	p := this.root
	for i, n := 0, len(word); i < n; i++ {
		if next, ok := p.next[word[i]]; !ok {
			return false
		} else {
			p = next
		}
	}

	return p.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	p := this.root
	for i, n := 0, len(prefix); i < n; i++ {
		if next, ok := p.next[prefix[i]]; !ok {
			return false
		} else {
			p = next
		}
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
