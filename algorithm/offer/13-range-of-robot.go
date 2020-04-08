package offer

/**
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？



示例 1：

输入：m = 2, n = 3, k = 1
输出：3
示例 1：

输入：m = 3, n = 1, k = 0
输出：1
提示：

1 <= n,m <= 100
0 <= k <= 20

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func getDigitSum(x int) int {
	res := 0
	for ; x > 0; x /= 10 {
		res += x % 10
	}
	return res
}

func isValidLocation(x, y, k int) bool {
	return getDigitSum(x)+getDigitSum(y) <= k
}

func movingCount(m int, n int, k int) int {
	q := []int{0}
	res := 1
	visited := make([]bool, m*n)
	visited[0] = true
	for len(q) > 0 {
		num := len(q)
		for i := 0; i < num; i++ {
			x, y := q[i]/n, q[i]%n
			if x+1 < m && !visited[q[i]+n] {
				visited[q[i]+n] = true
				if isValidLocation(x+1, y, k) {
					res++
					q = append(q, q[i]+n)
				}
			}
			if x > 0 && !visited[q[i]-n] {
				visited[q[i]-n] = true
				if isValidLocation(x-1, y, k) {
					res++
					q = append(q, q[i]-n)
				}
			}
			if y+1 < n && !visited[q[i]+1] {
				visited[q[i]+1] = true
				if isValidLocation(x, y+1, k) {
					res++
					q = append(q, q[i]+1)
				}
			}
			if y > 0 && !visited[q[i]-1] {
				visited[q[i]-1] = true
				if isValidLocation(x, y-1, k) {
					res++
					q = append(q, q[i]-1)
				}
			}
		}
		q = q[num:]
	}

	return res
}
