package leetcode

/**
The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.



Given an integer n, return all distinct solutions to the n-queens puzzle.

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space respectively.

Example:

Input: 4
Output: [
 [".Q..",  // Solution 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // Solution 2
  "Q...",
  "...Q",
  ".Q.."]
]
Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/n-queens
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func genChessBoardOutput(chess [][]bool) []string {
	board := []string{}
	for _, row := range chess {
		str := ""
		for _, e := range row {
			if e {
				str += "Q"
			} else {
				str += "."
			}
		}
		board = append(board, str)
	}
	return board
}

func isValidSolution(board [][]bool, n, i, j int) bool {
	sum, diff := i+j, i-j

	var x, y int

	for x, y = i, 0; y < j; y++ {
		if board[x][y] {
			return false
		}
	}

	for x = n - 1; x > i; x-- {
		y = sum - x
		if 0 <= y && y < n && board[x][sum-x] {
			return false
		}
	}

	for x = 0; x < i; x++ {
		y = x - diff
		if 0 <= y && y < n && board[x][y] {
			return false
		}
	}

	return true
}

func genChessBoard(boards *[][]string, board [][]bool, n, i, j int) {
	if j == n {
		*boards = append(*boards, genChessBoardOutput(board))
		return
	}

	for k := i; k < n; k++ {
		if !isValidSolution(board, n, k, j) {
			continue
		}
		board[k][j] = true
		genChessBoard(boards, board, n, 0, j+1)
		board[k][j] = false
	}
}

func solveNQueens(n int) [][]string {
	boardsOutput := [][]string{}

	board := make([][]bool, n, n)
	for i := range board {
		board[i] = make([]bool, n, n)
	}

	genChessBoard(&boardsOutput, board, n, 0, 0)

	return boardsOutput
}
