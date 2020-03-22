package backtracking

/**
Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.

Example:

Input: n = 4, k = 2
Output:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combinations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func combine(n int, k int) [][]int {
	res := [][]int{}
	combineSolution(&res, &[]int{}, n, 1, k)
	return res
}

func combineSolution(res *[][]int, c *[]int, n, i, k int) {
	if k == 0 {
		temp := make([]int, len(*c), len(*c))
		copy(temp, *c)
		*res = append(*res, temp)
		return
	}

	for ; i <= n; i++ {
		*c = append(*c, i)
		combineSolution(res, c, n, i+1, k-1)
		*c = (*c)[:len(*c)-1]
	}
}
