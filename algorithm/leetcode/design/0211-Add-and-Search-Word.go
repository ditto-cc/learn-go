package design

/**
Design a data structure that supports the following two operations:

void addWord(word)
bool search(word)
search(word) can search a literal word or a regular expression string containing only letters a-z or .. A . means it can represent any one letter.

Example:

addWord("bad")
addWord("dad")
addWord("mad")
search("pad") -> false
search("bad") -> true
search(".ad") -> true
search("b..") -> true
Note:
You may assume that all words are consist of lowercase letters a-z.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-and-search-word-data-structure-design
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
type trieNode struct {
	next   []*trieNode
	isWord bool
}

func newTrieNode() *trieNode {
	return &trieNode{isWord: false, next: make([]*trieNode, 26)}
}

type WordDictionary struct {
	root *trieNode
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{root: newTrieNode()}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	p := this.root
	for i, n := 0, len(word); i < n; i++ {
		index := word[i] - 'a'
		if p.next[index] == nil {
			p.next[index] = newTrieNode()
		}
		p = p.next[index]
	}
	p.isWord = true
}

func (node *trieNode) search(word *string, i, n int) bool {
	if node == nil {
		return false
	}

	if i >= n {
		return node.isWord
	}

	if (*word)[i] == '.' {
		for _, next := range node.next {
			if next.search(word, i+1, n) {
				return true
			}
		}
		return false
	} else {
		return node.next[(*word)[i]-'a'].search(word, i+1, n)
	}
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return this.root.search(&word, 0, len(word))
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
