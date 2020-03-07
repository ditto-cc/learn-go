package leetcode

/**
On a N x N grid of cells, each cell (x, y) with 0 <= x < N and 0 <= y < N has a lamp.

Initially, some number of lamps are on.  lamps[i] tells us the location of the i-th lamp that is on.  Each lamp that is on illuminates every square on its x-axis, y-axis, and both diagonals (similar to a Queen in chess).

For the i-th query queries[i] = (x, y), the answer to the query is 1 if the cell (x, y) is illuminated, else 0.

After each query (x, y) [in the order given by queries], we turn off any lamps that are at cell (x, y) or are adjacent 8-directionally (ie., share a corner or edge with cell (x, y).)

Return an array of answers.  Each value answer[i] should be equal to the answer of the i-th query queries[i].



Example 1:

Input: N = 5, lamps = [[0,0],[4,4]], queries = [[1,1],[1,0]]
Output: [1,0]
Explanation:
Before performing the first query we have both lamps [0,0] and [4,4] on.
The grid representing which cells are lit looks like this, where [0,0] is the top left corner, and [4,4] is the bottom right corner:
1 1 1 1 1
1 1 0 0 1
1 0 1 0 1
1 0 0 1 1
1 1 1 1 1
Then the query at [1, 1] returns 1 because the cell is lit.  After this query, the lamp at [0, 0] turns off, and the grid now looks like this:
1 0 0 0 1
0 1 0 0 1
0 0 1 0 1
0 0 0 1 1
1 1 1 1 1
Before performing the second query we have only the lamp [4,4] on.  Now the query at [1,0] returns 0, because the cell is no longer lit.


Note:

1 <= N <= 10^9
0 <= lamps.length <= 20000
0 <= queries.length <= 20000
lamps[i].length == queries[i].length == 2

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/grid-illumination
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func doMapPlus(m map[int]int, key, val int) {
	if _, ok := m[key]; ok {
		m[key] += val
		if m[key] == 0 {
			delete(m, key)
		}
	} else if val > 0 {
		m[key] = val
	}
}

func isLighted(m map[int]int, key int) bool {
	_, ok := m[key]
	return ok
}

func gridIllumination(N int, lamps [][]int, queries [][]int) []int {
	lampSet := map[int]bool{}

	rowMemo, colMemo := make(map[int]int), make(map[int]int)
	subDiaMemo, mainDiaMemo := make(map[int]int), make(map[int]int)

	for _, location := range lamps {
		doMapPlus(mainDiaMemo, location[0]-location[1]+N-1, 1)
		doMapPlus(subDiaMemo, location[0]+location[1], 1)
		doMapPlus(rowMemo, location[0], 1)
		doMapPlus(colMemo, location[1], 1)
		lampSet[N*location[0]+location[1]] = true
	}

	res := make([]int, len(queries), len(queries))
	for k, query := range queries {
		x, y := query[0], query[1]

		if isLighted(mainDiaMemo, x-y+N-1) || isLighted(subDiaMemo, x+y) ||
			isLighted(rowMemo, x) || isLighted(colMemo, y) {
			res[k] = 1
		}

		for i := x - 1; i <= x+1 && i < N; i++ {
			if i < 0 {
				continue
			}

			for j := y - 1; j <= y+1 && j < N; j++ {
				if j < 0 {
					continue
				}
				if _, ok := lampSet[N*i+j]; ok {
					doMapPlus(mainDiaMemo, i-j+N-1, -1)
					doMapPlus(subDiaMemo, i+j, -1)
					doMapPlus(rowMemo, i, -1)
					doMapPlus(colMemo, j, -1)

					delete(lampSet, i*N+j)
				}
			}
		}
	}

	return res

}
