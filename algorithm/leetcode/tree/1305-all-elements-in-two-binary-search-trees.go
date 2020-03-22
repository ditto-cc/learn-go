package tree

/**
Given two binary search trees root1 and root2.

Return a list containing all the integers from both trees sorted in ascending order.



Example 1:


Input: root1 = [2,1,4], root2 = [1,0,3]
Output: [0,1,1,2,3,4]
Example 2:

Input: root1 = [0,-10,10], root2 = [5,1,7,0,2]
Output: [-10,0,0,1,2,5,7,10]
Example 3:

Input: root1 = [], root2 = [5,1,7,0,2]
Output: [0,1,2,5,7]
Example 4:

Input: root1 = [0,-10,10], root2 = []
Output: [-10,0,10]
Example 5:


Input: root1 = [1,null,8], root2 = [8,1]
Output: [1,1,8,8]


Constraints:

Each tree has at most 5000 nodes.
Each node's value is between [-10^5, 10^5].

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/all-elements-in-two-binary-search-trees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {

	s1, s2 := []*TreeNode{}, []*TreeNode{}
	p, q := root1, root2
	arr1, arr2 := []int{}, []int{}
	for len(s1) > 0 || p != nil {
		if p != nil {
			s1 = append(s1, p)
			p = p.Left
		} else {
			top := s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
			arr1 = append(arr1, top.Val)
			p = top.Right
		}
	}

	for len(s2) > 0 || q != nil {
		if q != nil {
			s2 = append(s2, q)
			q = q.Left
		} else {
			top := s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
			arr2 = append(arr2, top.Val)
			q = top.Right
		}
	}

	res := []int{}
	for len(arr1) > 0 || len(arr2) > 0 {
		if len(arr1) == 0 {
			res = append(res, arr2[0])
			arr2 = arr2[1:]
		} else if len(arr2) == 0 {
			res = append(res, arr1[0])
			arr1 = arr1[1:]
		} else if arr1[0] < arr2[0] {
			res = append(res, arr1[0])
			arr1 = arr1[1:]
		} else {
			res = append(res, arr2[0])
			arr2 = arr2[1:]
		}
	}
	return res
}
