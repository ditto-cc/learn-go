package array

import "sort"

/**
Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

Note:

The solution set must not contain duplicate quadruplets.

Example:

Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/4sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func fourSum(nums []int, target int) [][]int {
	n, res := len(nums), [][]int{}
	if n < 4 {
		return res
	}
	sort.Ints(nums)
	for k, e := range nums {
		if k > 0 && nums[k-1] == e {
			continue
		}
		for l := k + 1; l < n; l++ {
			if l > k+1 && nums[l-1] == nums[l] {
				continue
			}
			for i, j := l+1, n-1; i < j; {
				val := nums[i] + nums[j] + nums[l] + e
				if val == target {
					res = append(res, []int{e, nums[l], nums[i], nums[j]})
					i++
					j--
					for ; i < j && nums[i] == nums[i-1]; i++ {
					}
					for ; i < j && nums[j] == nums[j+1]; j-- {
					}
				} else if val < target {
					i++
					for ; i < j && nums[i] == nums[i-1]; i++ {
					}
				} else {
					j--
					for ; i < j && nums[j] == nums[j+1]; j-- {
					}
				}
			}
		}
	}

	return res
}
