package backtracking

/**
Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

Each number in candidates may only be used once in the combination.

Note:

All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
Example 1:

Input: candidates = [10,1,2,7,6,1,5], target = 8,
A solution set is:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
Example 2:

Input: candidates = [2,5,2,1,2], target = 5,
A solution set is:
[
  [1,2,2],
  [5]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	genCombinationSum2(&res, candidates, []int{}, len(candidates)-1, target)
	return res
}

func genCombinationSum2(res *[][]int, candidates, comb []int, index, target int) {
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

		if i < index && candidates[i] == candidates[i+1] {
			continue
		}

		comb = append(comb, candidates[i])
		genCombinationSum2(res, candidates, comb, i-1, target-candidates[i])
		comb = comb[:len(comb)-1]
	}
}
