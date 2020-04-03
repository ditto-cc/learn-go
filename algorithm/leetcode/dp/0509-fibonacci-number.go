package dp

/**
The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,

F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), for N > 1.
Given N, calculate F(N).



Example 1:

Input: 2
Output: 1
Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.
Example 2:

Input: 3
Output: 2
Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.
Example 3:

Input: 4
Output: 3
Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.


Note:

0 ≤ N ≤ 30.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fibonacci-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func fib(N int) int {
	//pre, cur, next := 0, 1, 1
	//for ; N > 0; N-- {
	//	next = pre + cur
	//	pre = cur
	//	cur = next
	//}
	//return pre

	res, matrix := [][]int{{1, 0}, {0, 1}}, [][]int{{1, 1}, {1, 0}}
	for N--; N > 0; N >>= 1 {
		if N&1 == 1 {
			res = matrixProduct(res, matrix)
		}
		matrix = matrixProduct(matrix, matrix)
	}
	return res[0][0]
}

func matrixProduct(a, b [][]int) [][]int {
	res := [][]int{{0, 0}, {0, 0}}
	res[0][0] = a[0][0]*b[0][0] + a[0][1]*b[1][0]
	res[0][1] = a[0][0]*b[0][1] + a[0][1]*b[1][1]
	res[1][0] = a[1][0]*b[0][0] + a[1][1]*b[1][0]
	res[1][1] = a[1][1]*b[1][1] + a[1][0]*b[0][1]
	return res
}
