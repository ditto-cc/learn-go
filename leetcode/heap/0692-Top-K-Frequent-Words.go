package heap

import (
	"container/heap"
)

/**
Given a non-empty list of words, return the k most frequent elements.

Your answer should be sorted by frequency from highest to lowest. If two words have the same frequency, then the word with the lower alphabetical order comes first.

Example 1:
Input: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
Output: ["i", "love"]
Explanation: "i" and "love" are the two most frequent words.
    Note that "i" comes before "love" due to a lower alphabetical order.
Example 2:
Input: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k = 4
Output: ["the", "is", "sunny", "day"]
Explanation: "the", "is", "sunny" and "day" are the four most frequent words,
    with the number of occurrence being 4, 3, 2 and 1 respectively.
Note:
You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
Input words contain only lowercase letters.
Follow up:
Try to solve it in O(n log k) time and O(n) extra space.
*/
type myHeap struct {
	data    []string
	counter map[string]int
}

func (h *myHeap) Push(x interface{}) {
	h.data = append(h.data, x.(string))
}

func (h *myHeap) Pop() interface{} {
	ret := h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]
	return ret
}

func (h *myHeap) Len() int {
	return len(h.data)
}

func (h *myHeap) Less(i, j int) bool {
	if h.counter[h.data[i]] < h.counter[h.data[j]] {
		return true
	}

	if h.counter[h.data[i]] == h.counter[h.data[j]] {
		return h.data[i] > h.data[j]
	}

	return false
}

func (h *myHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func topKFrequent(words []string, k int) []string {
	h := myHeap{
		data:    make([]string, 0, k+1),
		counter: make(map[string]int),
	}

	for _, s := range words {
		if _, ok := h.counter[s]; ok {
			h.counter[s]++
		} else {
			h.counter[s] = 1
		}
	}

	for str := range h.counter {
		heap.Push(&h, str)
		if h.Len() > k {
			heap.Pop(&h)
		}
	}

	res := make([]string, k, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(&h).(string)
	}

	return res
}
