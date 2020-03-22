package dp

/**
On a staircase, the i-th step has some non-negative cost cost[i] assigned (0 indexed).

Once you pay the cost, you can either climb one or two steps. You need to find minimum cost to reach the top of the floor, and you can either start from the step with index 0, or the step with index 1.

Example 1:
Input: cost = [10, 15, 20]
Output: 15
Explanation: Cheapest is start on cost[1], pay that cost and go to the top.
Example 2:
Input: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
Output: 6
Explanation: Cheapest is start on cost[0], and only step on 1s, skipping cost[3].
Note:
cost will have a length in the range [2, 1000].
Every cost[i] will be an integer in the range [0, 999].

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/min-cost-climbing-stairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func minCostClimbingStairs(cost []int) int {
	pre, cur, next := cost[0], cost[1], 0
	cost = append(cost, 0)
	n := len(cost)
	for i := 2; i < n; i++ {
		next = min(pre+cost[i], cur+cost[i])
		pre = cur
		cur = next
	}
	return cur
}
