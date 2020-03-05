package array

import (
	"math"
	"sort"
)

/**
Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

Example:

Given array nums = [-1, 2, 1, -4], and target = 1.

The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum-closest
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func abs(val int) int {
	if val < 0 {
		return 0 - val
	}
	return val
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	res, diff := 0, math.MaxInt32
	for k, e := range nums {
		for i, j := k+1, len(nums)-1; i < j; {
			val := nums[i] + nums[j] + e
			if val == target {
				return val
			} else if val < target {
				i++
				for ; i < j && nums[i] == nums[i-1]; i++ {
				}
			} else {
				j--
				for ; i < j && nums[j] == nums[j+1]; j-- {
				}
			}
			newDiff := abs(val - target)
			if newDiff < diff {
				diff = newDiff
				res = val
			}
		}
	}
	return res
}
