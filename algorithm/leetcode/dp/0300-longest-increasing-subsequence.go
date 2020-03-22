package dp

/**
Given an unsorted array of integers, find the length of longest increasing subsequence.

Example:

Input: [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
Note:

There may be more than one LIS combination, it is only necessary for you to return the length.
Your algorithm should run in O(n2) complexity.
Follow up: Could you improve it to O(n log n) time complexity?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-increasing-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// DP
// O(N^2)
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	res := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] >= nums[i] {
				continue
			}
			dp[i] = max(dp[i], 1+dp[j])
			res = max(dp[i], res)
		}
	}

	return res
}

// Followed Up
// BinSearch
// O(NlogN)

// func lengthOfLIS(nums []int) int {
// 	subArray := []int{}
// 	l := 0
// 	for _, e := range nums {
// 		i := sort.SearchInts(subArray, e)
// 		if i < l {
// 			subArray[i] = e
// 		} else {
// 			subArray = append(subArray, e)
// 			l++
// 		}
// 	}

// 	return l
// }
