package sort

/**
Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

Note: You are not suppose to use the library's sort function for this problem.

Example:

Input: [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
Follow up:

A rather straight forward solution is a two-pass algorithm using counting sort.
First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.
Could you come up with a one-pass algorithm using only constant space?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-colors
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func quickSort(nums []int, l, r int) {
	
	if l >= r {
		return
	}

	lt, gt, i := l, r + 1, l+1
	for i < gt {
		if nums[i] == nums[l] {
			i++
		} else if nums[i] < nums[l] {
			nums[i], nums[lt+1] = nums[lt+1], nums[i]
			lt++
			i++
		} else {
			nums[i], nums[gt-1] = nums[gt-1], nums[i]
			gt--
		}
	}
	nums[l], nums[lt] = nums[lt], nums[l]

	quickSort(nums, l, lt-1)
	quickSort(nums, gt, r)
}

func sortColors(nums []int)  {
	quickSort(nums, 0, len(nums)-1)
}