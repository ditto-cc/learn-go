package array

/**
Given a collection of distinct integers, return all possible permutations.

Example:

Input: [1,2,3]
Output:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func permute(nums []int) [][]int {
	res := [][]int{}

	genPermute(&res, nums, 0, len(nums))

	return res
}

func genPermute(res *[][]int, nums []int, index, n int) {
	if index == n {
		temp := make([]int, n)
		copy(temp, nums)
		*res = append((*res), temp)
		return
	}

	for i := index; i < n; i++ {
		nums[i], nums[index] = nums[index], nums[i]
		genPermute(res, nums, index+1, n)
		nums[i], nums[index] = nums[index], nums[i]
	}

}
