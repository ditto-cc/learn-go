package array

/**
Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

Example 1:

Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Example 2:

Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/plus-one
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func plusOne(digits []int) []int {
	n, carry := len(digits), 0
	for i := n - 1; i >= 0 && carry == 1; i-- {
		r := digits[i] + carry
		carry = r / 10
		digits[i] = r % 10
	}

	if carry == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}
