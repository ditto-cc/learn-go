package heap

/**
Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Example 1:

Input: [3,2,1,5,6,4] and k = 2
Output: 5
Example 2:

Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
Note:
You may assume k is always valid, 1 ≤ k ≤ array's length.
*/

func down(h []int, i, end int) {
	for 2*i+1 < end {
		c := 2*i + 1
		if c+1 < end && h[c+1] < h[c] {
			c++
		}

		if h[i] < h[c] {
			break
		}

		h[i], h[c] = h[c], h[i]
		i = c
	}
}

func findKthLargest(nums []int, k int) int {
	for i := (k - 1) / 2; i >= 0; i-- {
		down(nums, i, k)
	}

	for i := k; i < len(nums); i++ {
		if nums[i] > nums[0] {
			nums[0] = nums[i]
			down(nums, 0, k)
		}
	}

	return nums[0]
}
