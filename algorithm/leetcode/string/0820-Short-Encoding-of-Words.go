package string

/**
 * Given a list of words, we may encode it by writing a reference string S and a list of indexes A.
 *
 * For example, if the list of words is ["time", "me", "bell"], we can write it as S = "time#bell#" and indexes = [0, 2, 5].
 *
 * Then for each index, we will recover the word by reading from the reference string from that index until we reach a "#" character.
 *
 * What is the length of the shortest reference string S possible that encodes the given words?
 *
 * Example:
 *
 * Input: words = ["time", "me", "bell"]
 * Output: 10
 * Explanation: S = "time#bell#" and indexes = [0, 2, 5].
 *
 *
 * Note:
 *
 * 1 <= words.length <= 2000.
 * 1 <= words[i].length <= 7.
 * Each word has only lowercase letters.
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/short-encoding-of-words
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

type trieNode struct {
	m      []*trieNode
	isWord bool
}

func minimumLengthEncoding(words []string) int {
	if len(words) == 0 {
		return 0
	}

	res := 0
	trie := &trieNode{m: make([]*trieNode, 26)}
	for _, word := range words {
		n := len(word)
		flag, p := false, trie
		for i := n - 1; i >= 0; i-- {
			index := word[i] - 'a'
			if p.m[index] == nil {
				p.m[index] = &trieNode{m: make([]*trieNode, 26)}
				flag = true
				if p.isWord {
					res -= n - i
					p.isWord = false
				}
			}
			p = p.m[index]
		}

		if flag {
			p.isWord = true
			res += n + 1
		}
	}

	if len(trie.m) == 0 {
		return 1
	}

	return res

	// Use Map Store Length of Words
	// Delete All Suffix
	// Then Sum Words Length Together

	//m := map[string]int{}
	//
	//for _, word := range words {
	//	m[word] = len(word)
	//}
	//
	//for _, word := range words {
	//	n := len(word)
	//	for i := 1; i < n; i++ {
	//		if _, ok := m[word[i:]]; ok {
	//			delete(m, word[i:])
	//		}
	//	}
	//}
	//
	//var res int
	//for _, l := range m {
	//	res += l + 1
	//}
	//return res
}
