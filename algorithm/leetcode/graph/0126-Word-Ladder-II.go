package graph

/*
给定两个单词（beginWord 和 endWord）和一个字典 wordList，找出所有从 beginWord 到 endWord 的最短转换序列。转换需遵循如下规则：

每次转换只能改变一个字母。
转换过程中的中间单词必须是字典中的单词。
说明:

如果不存在这样的转换序列，返回一个空列表。
所有单词具有相同的长度。
所有单词只由小写字母组成。
字典中不存在重复的单词。
你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
示例 1:

输入:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

输出:
[
  ["hit","hot","dot","dog","cog"],
  ["hit","hot","lot","log","cog"]
]
示例 2:

输入:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

输出: []

解释: endWord "cog" 不在字典中，所以不存在符合要求的转换序列。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-ladder-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

type trieNode struct {
	next  map[byte]*trieNode
	index int
}

func (trie *trieNode) add(word string, index, wordLength int) {
	root := trie
	for i := 0; i < wordLength; i++ {
		if next, ok := root.next[word[i]]; !ok {
			next = &trieNode{next: map[byte]*trieNode{}, index: -1}
			root.next[word[i]] = next
			root = next
		} else {
			root = next
		}
	}
	root.index = index
}

func (trie *trieNode) search(word string, n int) int {
	node := trie
	for i := 0; i < n; i++ {
		next, ok := node.next[word[i]]
		if !ok {
			return -1
		}
		node = next
	}
	return node.index
}

func (trie *trieNode) searchPattern(word string, i, dot, wordLength int, res *[]int) {
	if trie == nil {
		return
	}

	if i == wordLength {
		if trie.index != -1 {
			*res = append(*res, trie.index)
		}
		return
	}

	if i == dot {
		for j, node := range trie.next {
			if j == word[i] {
				continue
			}
			node.searchPattern(word, i+1, dot, wordLength, res)
		}
	} else {
		trie.next[word[i]].searchPattern(word, i+1, dot, wordLength, res)
	}
}

func bfs(adjList [][]int, start, end int) [][]int {
	parent := make([][]int, len(adjList))
	q := []int{start}
	parent[start] = append(parent[start], -1)
	visit := make([]bool, len(adjList))
	visit[start] = true
	hasEnd := false
	for len(q) > 0 && !hasEnd {
		size := len(q)
		temp := map[int]bool{}
		for i := 0; i < size; i++ {
			for _, adj := range adjList[q[i]] {
				hasEnd = hasEnd || adj == end
				if !visit[adj] {
					parent[adj] = append(parent[adj], q[i])
					temp[adj] = true
				}
			}
		}

		for adj := range temp {
			visit[adj] = true
			q = append(q, adj)
		}

		q = q[size:]
	}
	return parent
}

func generatePath(parents [][]int, paths *[][]int, path *[]int, i int) {
	if i == -1 {
		temp := make([]int, len(*path))
		copy(temp, *path)
		*paths = append(*paths, temp)
		return
	}

	*path = append(*path, i)
	for _, parent := range parents[i] {
		generatePath(parents, paths, path, parent)
	}
	*path = (*path)[:len(*path)-1]
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	trie := &trieNode{next: map[byte]*trieNode{}, index: -1}
	wordLength := len(beginWord)
	// 构建Trie
	for i, word := range wordList {
		// 添加wordList里的字符串
		trie.add(word, i, wordLength)
	}

	// endWord不在Trie中，直接返回nil
	endIndex := -1
	if endIndex = trie.search(endWord, wordLength); endIndex == -1 {
		return nil
	}

	// 计算开始节点下标
	// 若wordList中不存在beginWord，添加beginWord到wordList末尾
	// 若存在，由Trie得到开始下标startIndex
	startIndex := -1
	if startIndex = trie.search(beginWord, wordLength); startIndex == -1 {
		wordList = append(wordList, beginWord)
		startIndex = len(wordList) - 1
	}

	// 根据Trie计算每个单词只修改一个位置可以转换的下标，得到无向图的邻接表
	// O(len(word) * len(word) * n)
	adjList := make([][]int, len(wordList))
	for i, word := range wordList {
		for j := 0; j < wordLength; j++ {
			trie.searchPattern(word, 0, j, wordLength, &adjList[i])
		}
	}

	// 广度优先搜索，得到parents数组，记录每个下标的前驱（可能不唯一）
	var paths [][]int
	parents := bfs(adjList, startIndex, endIndex)
	// 得到索引路径的逆序
	generatePath(parents, &paths, &[]int{}, endIndex)
	if len(paths) == 0 {
		return nil
	}

	// 得到最终结果
	var res [][]string
	n := len(paths[0])
	for _, path := range paths {
		temp := make([]string, n)
		for i := n - 1; i >= 0; i-- {
			temp[n - i - 1] = wordList[path[i]]
		}
		res = append(res, temp)
	}

	return res
}