package search

/**
In a given grid, each cell can have one of three values:

the value 0 representing an empty cell;
the value 1 representing a fresh orange;
the value 2 representing a rotten orange.
Every minute, any fresh orange that is adjacent (4-directionally) to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange.  If this is impossible, return -1 instead.



Example 1:



Input: [[2,1,1],[1,1,0],[0,1,1]]
Output: 4
Example 2:

Input: [[2,1,1],[0,1,1],[1,0,1]]
Output: -1
Explanation:  The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.
Example 3:

Input: [[0,2]]
Output: 0
Explanation:  Since there are already no fresh oranges at minute 0, the answer is just 0.


Note:

1 <= grid.length <= 10
1 <= grid[0].length <= 10
grid[i][j] is only 0, 1, or 2.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotting-oranges
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func orangesRotting(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	q := [][]int{}
	hasOrange := false
	for i, row := range grid {
		for j, e := range row {
			if e == 2 {
				q = append(q, []int{i, j})
			} else if !hasOrange && e == 1 {
				hasOrange = true
			}
		}
	}

	hasRotten := len(q) > 0
	if !hasRotten && hasOrange {
		return -1
	} else if (hasRotten && !hasOrange) || (!hasRotten && !hasOrange) {
		return 0
	}

	res := -1
	for len(q) > 0 {
		res++
		num := len(q)
		for ; num > 0; num-- {
			i, j := q[0][0], q[0][1]
			q = q[1:]
			if i-1 >= 0 && grid[i-1][j] == 1 {
				grid[i-1][j] = 2
				q = append(q, []int{i - 1, j})
			}
			if j-1 >= 0 && grid[i][j-1] == 1 {
				grid[i][j-1] = 2
				q = append(q, []int{i, j - 1})
			}
			if i+1 < n && grid[i+1][j] == 1 {
				grid[i+1][j] = 2
				q = append(q, []int{i + 1, j})
			}
			if j+1 < m && grid[i][j+1] == 1 {
				grid[i][j+1] = 2
				q = append(q, []int{i, j + 1})
			}
		}
	}

	for _, row := range grid {
		for _, e := range row {
			if e == 1 {
				return -1
			}
		}
	}

	return res
}
