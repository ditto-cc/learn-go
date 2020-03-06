package leetcode

import "strconv"

/**
The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

Given two integers x and y, calculate the Hamming distance.

Note:
0 ≤ x, y < 231.

Example:

Input: x = 1, y = 4

Output: 2

Explanation:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑

The above arrows point to positions where the corresponding bits are different.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/hamming-distance
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func hammingDistance(x int, y int) int {
	a, b := strconv.FormatInt(int64(x), 2), strconv.FormatInt(int64(y), 2)
	if x < y {
		a, b = b, a
	}
	al := len(a)
	bl := len(b)
	res := 0
	var i, j int
	for i = 0; i < al - bl; i++ {
		if a[i] == '1' {
			res ++
		}
	}
	for i < al && j < bl {
		if a[i] != b[j] {
			res ++
		}
		i++
		j++
	}
	return res
}
