package leetcode

import "sort"

/**
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	res := [][]int{}
	for k, target := range nums {
		if nums[k] > 0 {
			break
		}
		if k > 0 && nums[k-1] == target {
			continue
		}
		for i, j := k+1, len(nums)-1; i < j; {
			val := nums[i] + nums[j] + target
			if val == 0 {
				res = append(res, []int{target, nums[i], nums[j]})
				i++
				j--
				for ; i < j && nums[i] == nums[i-1]; i++ {
				}
				for ; i < j && nums[j] == nums[j+1]; j-- {
				}
			} else if val < 0 {
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
	return res
}
