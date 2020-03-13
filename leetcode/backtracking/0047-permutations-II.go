package backtracking

import "sort"

/**
Given a collection of numbers that might contain duplicates, return all possible unique permutations.

Example:

Input: [1,1,2]
Output:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	n := len(nums)
	visit := make([]bool, n)

	genPermute(&res, []int{}, nums, visit, n)

	return res
}

func genPermute(res *[][]int, per, nums []int, visit []bool, remain int) {

	if remain == 0 {
		temp := make([]int, len(per))
		copy(temp, per)
		*res = append(*res, temp)
		return
	}

	for i := range nums {

		if visit[i] || (i > 0 && nums[i] == nums[i-1] && !visit[i-1]) {
			continue
		}

		visit[i] = true
		per = append(per, nums[i])

		genPermute(res, per, nums, visit, remain-1)

		visit[i] = false
		per = per[:len(per)-1]
	}

}
