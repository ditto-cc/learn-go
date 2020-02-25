package leetcode

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

func fourSum(nums []int, target int) [][]int {
	var data MyType = nums
	n, res := data.Len(), [][]int{}
	sort.Sort(data)
	for k, e := range nums {
		for l := k + 1; l < n; l++ {
			for i, j := l+1, n-1; i < j; {
				val := data[i] + data[j] + data[l] + e
				if val == target {
					res = append(res, []int{e, data[l], data[i], data[j]})
					i++
					j--
					for ; i < j && data[i] == data[i-1]; i++ {
					}
					for ; i < j && data[j] == data[j+1]; j-- {
					}
				} else if val < target {
					i++
					for ; i < j && data[i] == data[i-1]; i++ {
					}
				} else {
					j--
					for ; i < j && data[j] == data[j+1]; j-- {
					}
				}
			}
		}
	}

	return res
}
