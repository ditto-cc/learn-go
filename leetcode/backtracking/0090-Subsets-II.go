package backtracking

import "sort"

/**
Given a collection of integers that might contain duplicates, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:

Input: [1,2,2]
Output:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	visit := make([]bool, len(nums))

	sort.Ints(nums)

	subsetsWithDupSolution(&res, nums, []int{}, visit, 0, len(nums))

	return res
}

func subsetsWithDupSolution(res *[][]int, nums, set []int, visit []bool, i, n int) {
	if i < -1 {
		return
	}

	temp := make([]int, len(set))
	copy(temp, set)
	*res = append(*res, temp)

	for ; i < n; i++ {
		if visit[i] || (i > 0 && nums[i-1] == nums[i] && !visit[i-1]) {
			continue
		}
		set = append(set, nums[i])
		subsetsWithDupSolution(res, nums, set, visit, i+1, n)
		set = set[:len(set)-1]
	}
}
