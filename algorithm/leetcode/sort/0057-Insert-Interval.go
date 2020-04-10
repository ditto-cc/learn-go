package sort

/*
Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).

You may assume that the intervals were initially sorted according to their start times.

Example 1:

Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]
Example 2:

Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/insert-interval
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	n := len(intervals)
	i := 0

	// 添加小于NewInterval的区间
	for ; i < n; i++ {
		if intervals[i][0] < newInterval[0] {
			res = append(res, intervals[i])
		} else {
			break
		}
	}

	// 合并res最后一个区间和NewInterval区间
	last := i - 1
	if newInterval[0] <= res[last][1] {
		res[last][1] = max(res[last][1], newInterval[1])
	} else {
		res = append(res, newInterval)
		last++
	}

	// 合并res最后一个区间和原数组区间
	for ; i < n; i++ {
		if intervals[i][0] <= res[last][1] {
			res[last][1] = max(res[last][1], intervals[i][1])
		} else {
			break
		}
	}

	// 添加之后的不相交区间
	for ; i < n; i++ {
		res = append(res, intervals[i])
	}

	return res
}
