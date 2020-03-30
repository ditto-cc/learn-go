package sort

/**
 * Given an array of integers nums, sort the array in ascending order.
 *
 *
 *
 * Example 1:
 *
 * Input: nums = [5,2,3,1]
 * Output: [1,2,3,5]
 * Example 2:
 *
 * Input: nums = [5,1,1,2,0,0]
 * Output: [0,0,1,1,2,5]
 *
 *
 * Constraints:
 *
 * 1 <= nums.length <= 50000
 * -50000 <= nums[i] <= 50000
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/sort-an-array
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func sortArray(nums []int) []int {
	qSort(nums, 0, len(nums))
	return nums
}

func qSort(nums []int, l, r int) {
	if l >= r {
		return
	}

	lt, gt := l, r
	for i := l + 1; i < gt; {
		if nums[i] == nums[l] {
			i++
		} else if nums[i] < nums[l] {
			nums[lt+1], nums[i] = nums[i], nums[lt+1]
			lt++
			i++
		} else {
			nums[i], nums[gt-1] = nums[gt-1], nums[i]
			gt--
		}
	}
	nums[l], nums[lt] = nums[lt], nums[l]

	qSort(nums, l, lt)
	qSort(nums, gt, r)
}
