package dp

/**
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// f(0) = nums[0]
// f(k) = max{ f(k-1) + nums[k], nums[k] }

func maxSubArray(nums []int) int {
	res, maxSoFar, n := nums[0], nums[0], len(nums)
	for i := 1; i < n; i++ {

		maxSoFar = max(nums[i], nums[i]+maxSoFar)
		res = max(res, maxSoFar)
	}
	return res
}
