package leetcode

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
type MyType []int

func (data MyType) Len() int {
	return len(data)
}
func (data MyType) Less(i, j int) bool {
	if data[i] <= data[j] {
		return true
	}
	return false
}

// Swap swaps the elements with indexes i and j.
func (data MyType) Swap(i, j int) {
	data[i], data[j] = data[j], data[i]
}

func abs(val int) int {
	if val < 0 {
		return 0 - val
	}
	return val
}

func threeSumClosest(nums []int, target int) int {
	var data MyType = nums
	sort.Sort(data)

	res, diff := 0, math.MaxInt32
	for k, e := range data {
		if nums[k] > 0 {
			break
		}
		if k > 0 && nums[k-1] == e {
			continue
		}
		for i, j := k+1, data.Len()-1; i < j; {
			val := data[i] + data[j] + e
			if val == target {
				return val
			} else if val < target {
				i++
				for ; i < j && data[i] == data[i-1]; i++ {
				}
			} else {
				j--
				for ; i < j && data[j] == data[j+1]; j-- {
				}
			}
			if abs(val-diff) < diff {
				diff = abs(val - diff)
				res = val
			}
		}
	}
	return res
}
