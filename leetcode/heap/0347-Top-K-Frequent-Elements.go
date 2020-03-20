package heap

/**
Given a non-empty array of integers, return the k most frequent elements.

Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
Example 2:

Input: nums = [1], k = 1
Output: [1]
Note:

You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
*/

func down(data []int, i, end int, freMap map[int]int) {
	for 2*i+1 < end {
		c := 2*i + 1
		if c+1 < end && freMap[data[c+1]] < freMap[data[c]] {
			c++
		}

		if freMap[data[i]] <= freMap[data[c]] {
			break
		}

		data[i], data[c] = data[c], data[i]
		i = c
	}
}

func topKFrequent(nums []int, k int) []int {
	if k == 0 {
		return nil
	}

	freMap := map[int]int{}
	for _, e := range nums {
		if _, ok := freMap[e]; ok {
			freMap[e]++
		} else {
			freMap[e] = 1
		}
	}

	h := make([]int, 0, len(freMap))
	for i := range freMap {
		h = append(h, i)
	}

	for i := (k - 1) / 2; i >= 0; i-- {
		down(h, i, k, freMap)
	}

	for i := k; i < len(h); i++ {
		if freMap[h[i]] > freMap[h[0]] {
			h[0] = h[i]
			down(h, 0, k, freMap)
		}
	}

	return h[:k]
}
