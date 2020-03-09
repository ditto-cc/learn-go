package search

/**
Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func searchRange(nums []int, target int) []int {
	n, res := len(nums), []int{-1, -1}

	l, r := 0, n-1
	for l <= r {
		mid := int(uint(l+r) >> 1)
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	if l >= n || nums[l] != target {
		return res
	}

	res[0] = l

	r = n - 1
	for l <= r {
		mid := int(uint(l+r) >> 1)
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	res[1] = r

	return res
}
