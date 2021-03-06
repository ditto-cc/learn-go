package dp

/**
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.

Example 1:

Input: [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
             Total amount you can rob = 1 + 3 = 4.
Example 2:

Input: [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
             Total amount you can rob = 2 + 9 + 1 = 12.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func rob(nums []int) int {
	n := len(nums)
	switch {
	case 0 == n:
		return 0
	case 1 == n:
		return nums[0]
	case 2 < n:
		nums[2] += nums[0]
		for i := 3; i < n; i++ {
			nums[i] += max(nums[i-2], nums[i-3])
		}
	}
	return max(nums[n-1], nums[n-2])
}
