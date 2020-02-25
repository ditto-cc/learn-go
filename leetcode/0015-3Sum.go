package leetcode

import "sort"

/**
Given an array data of n integers, are there elements a, b, c in data such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array data = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
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

func threeSum(nums []int) [][]int {
	var data MyType = nums
	sort.Sort(data)

	res := [][]int{}
	for k, target := range data {
		if nums[k] > 0 {
			break
		}
		if k > 0 && nums[k-1] == target {
			continue
		}
		for i, j := k+1, data.Len()-1; i < j; {
			val := data[i] + data[j] + target
			if val == 0 {
				res = append(res, []int{target, data[i], data[j]})
				i++
				j--
				for ; i < j && data[i] == data[i-1]; i++ {
				}
				for ; i < j && data[j] == data[j+1]; j-- {
				}
			} else if val < 0 {
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
	return res
}
