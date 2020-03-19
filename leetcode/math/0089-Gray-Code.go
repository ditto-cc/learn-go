package math

/**
The gray code is a binary numeral system where two successive values differ in only one bit.

Given a non-negative integer n representing the total number of bits in the code, print the sequence of gray code. A gray code sequence must begin with 0.

Example 1:

Input: 2
Output: [0,1,3,2]
Explanation:
00 - 0
01 - 1
11 - 3
10 - 2

For a given n, a gray code sequence may not be uniquely defined.
For example, [0,2,3,1] is also a valid gray code sequence.

00 - 0
10 - 2
11 - 3
01 - 1
Example 2:

Input: 0
Output: [0]
Explanation: We define the gray code sequence to begin with 0.
             A gray code sequence of n has maxBit = 2n, which for n = 0 the maxBit is 20 = 1.
             Therefore, for n = 0 the gray code sequence is [0].
*/

func join(a, b int) int {
	if a&b > 0 {
		return 1
	}
	return 0
}

func convert(e, maxBit int) int {
	res := join(e, maxBit)
	pre, cur := res, 0
	for maxBit >>= 1; maxBit > 0; maxBit >>= 1 {
		cur = join(maxBit, e)
		res <<= 1
		res += cur ^ pre
		pre = cur
	}
	return res
}

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	size := 1 << n
	res := make([]int, size)
	for i := range res {
		res[i] = i
	}

	maxBit := size >> 1
	for i, e := range res {
		res[i] = convert(e, maxBit)
	}

	return res
}
