package sort

import "sort"

/**
 * Given a collection of intervals, merge all overlapping intervals.
 *
 * Example 1:
 *
 * Input: [[1,3],[2,6],[8,10],[15,18]]
 * Output: [[1,6],[8,10],[15,18]]
 * Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
 * Example 2:
 *
 * Input: [[1,4],[4,5]]
 * Output: [[1,5]]
 * Explanation: Intervals [1,4] and [4,5] are considered overlapping.
 * NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.
 */

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})

	res := [][]int{intervals[0]}
	last := 0
	for i := 1; i < n; i++ {
		if res[last][1] < intervals[i][0] {
			res = append(res, intervals[i])
			last++
		} else {
			res[last][1] = max(intervals[i][1], res[last][1])
		}
	}

	return res
}
