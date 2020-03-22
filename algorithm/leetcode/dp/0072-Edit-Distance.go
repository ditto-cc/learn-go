package dp

/**
 * Given two words word1 and word2, find the minimum number of operations required to convert word1 to word2.
 *
 * You have the following 3 operations permitted on a word:
 *
 * Insert a character
 * Delete a character
 * Replace a character
 * Example 1:
 *
 * Input: word1 = "horse", word2 = "ros"
 * Output: 3
 * Explanation:
 * horse -> rorse (replace 'h' with 'r')
 * rorse -> rose (remove 'r')
 * rose -> ros (remove 'e')
 * Example 2:
 *
 * Input: word1 = "intention", word2 = "execution"
 * Output: 5
 * Explanation:
 * intention -> inention (remove 't')
 * inention -> enention (replace 'i' with 'e')
 * enention -> exention (replace 'n' with 'x')
 * exention -> exection (replace 'n' with 'c')
 * exection -> execution (insert 'u')
 */

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}

	if b < c {
		return b
	}
	return c
}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m > n {
		return minDistanceSolution(&word1, &word2, m, n)
	}
	return minDistanceSolution(&word2, &word1, n, m)
}

func minDistanceSolution(l, s *string, ll, sl int) int {
	dp := make([][]int, 2)
	dp[0], dp[1] = make([]int, sl+1), make([]int, sl+1)

	for i := range dp[0] {
		dp[0][i] = i
	}

	for i, index := 1, 1; i <= ll; i, index = i+1, 1-index {
		dp[index][0] = i
		for j := 1; j <= sl; j++ {
			if (*l)[i-1] == (*s)[j-1] {
				dp[index][j] = dp[1-index][j-1]
			} else {
				dp[index][j] = 1 + min(dp[1-index][j], dp[index][j-1], dp[1-index][j-1])
			}
		}
	}

	return dp[ll&1][sl]
}
