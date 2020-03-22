package math

/**

Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only store integers
within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem,
assume that your function returns 0 when the reversed integer overflows.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func reverse(x int) int {
	res := 0
	for x != 0 {
		res *= 10
		res += x % 10
		x /= 10
	}

	if res > (1<<31)-1 || res < -((1<<31)-1) {
		return 0
	}

	return res
}
