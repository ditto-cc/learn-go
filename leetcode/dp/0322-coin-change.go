package leetcode

import (
	"math"
	"sort"
)

/**
You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

Example 1:

Input: coins = [1, 2, 5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
Example 2:

Input: coins = [2], amount = 3
Output: -1
Note:
You may assume that you have an infinite number of each kind of coin.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/coin-change
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	res := math.MaxInt32
	sort.Ints(coins)
	coinChangeSolution(coins, amount, 0, len(coins)-1, &res)

	if res != math.MaxInt32 {
		return res
	}
	return -1

}

func coinChangeSolution(coins []int, amount, count, index int, res *int) {
	if amount == 0 {
		if count < *res {
			*res = count
		}
		return
	}

	if index == -1 {
		return
	}

	for i := amount / coins[index]; i >= 0 && count+i < *res; i-- {
		coinChangeSolution(coins, amount-i*coins[index], count+i, index-1, res)
	}
}

//
//func coinChange(coins []int, amount int) int {
//	n, m := len(coins), amount + 1
//	memo := make([]int, m, m)
//	for i := 1; i < m; i++ {
//		memo[i] = m
//	}
//
//	for i := 1; i < m; i++ {
//		for j := 0; j < n; j++ {
//			remain := i - coins[j]
//			if remain < 0 {
//				continue
//			}
//
//			memo[i] = min(memo[i], memo[remain] + 1)
//		}
//	}
//
//	if memo[amount] == m {
//		return -1
//	}
//	return memo[amount]
//}
//
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
