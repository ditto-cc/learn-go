package math

/**
Find the nth digit of the infinite integer sequence 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...

Note:
n is positive and will fit within the range of a 32-bit signed integer (n < 231).

Example 1:

Input:
3

Output:
3
Example 2:

Input:
11

Output:
0

Explanation:
The 11th digit of the sequence 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... is a 0, which is part of the number 10.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/nth-digit
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findNthDigit(n int) int {
	pre, num, index := 0, 0, 0
	digit := 1
	for l, r := 0, 10; ; l, r, digit = r, r*10, digit+1 {
		if n-pre < (r-l)*digit {
			i := (n - pre) / digit
			num = l + i
			index = n - pre - digit*i
			break
		}
		pre += (r - l) * digit
	}

	res := 0
	for index = digit - index; index > 0; index-- {
		res = num % 10
		num /= 10
	}
	return res
}
