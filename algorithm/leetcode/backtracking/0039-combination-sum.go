package backtracking

/**
Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

The same repeated number may be chosen from candidates unlimited number of times.

Note:

All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
Example 1:

Input: candidates = [2,3,6,7], target = 7,
A solution set is:
[
  [7],
  [2,2,3]
]
Example 2:

Input: candidates = [2,3,5], target = 8,
A solution set is:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	genCombinationSum(&res, candidates, []int{}, len(candidates)-1, target)
	return res
}

func genCombinationSum(res *[][]int, candidates, comb []int, index, target int) {
	if target < 0 {
		return
	}

	if target == 0 {
		temp := make([]int, len(comb))
		copy(temp, comb)
		*res = append(*res, temp)
		return
	}

	for i := index; i >= 0; i-- {
		comb = append(comb, candidates[i])
		genCombinationSum(res, candidates, comb, i, target-candidates[i])
		comb = comb[:len(comb)-1]
	}
}
