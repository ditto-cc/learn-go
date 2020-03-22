package backtracking

/**
Given a set of distinct integers, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:

Input: nums = [1,2,3]
Output:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func subsets(nums []int) [][]int {
	var res [][]int
	subsetSolution(&res, nums, []int{}, len(nums)-1)
	return res
}

func subsetSolution(res *[][]int, nums, set []int, i int) {
	if i < -1 {
		return
	}

	temp := make([]int, len(set))
	copy(temp, set)
	*res = append(*res, temp)

	for j := i; j >= 0; j-- {
		set = append(set, nums[j])
		subsetSolution(res, nums, set, j-1)
		set = set[:len(set)-1]
	}
}
