package array

/**
Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.

Example 1:

Input: [3,2,3]
Output: 3
Example 2:

Input: [2,2,1,1,1,2,2]
Output: 2

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/majority-element
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func majorityElement(nums []int) int {
	res, count := nums[0], 1
	n := len(nums)
	for i := 1; i < n; i++ {
		if res == nums[i] {
			count ++
		} else {
			count --
		}
		if count == 0 {
			res = nums[i]
			count=1
		}
	}

    return res
}