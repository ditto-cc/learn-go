package uf

/*
给定一个未排序的整数数组，找出最长连续序列的长度。

要求算法的时间复杂度为 O(n)。

示例:

输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-consecutive-sequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func union(parent map[int][]int, p, q int, maxSize *int) {
	pRoot, qRoot := find(parent, p), find(parent, q)
	if pRoot == qRoot {
		return
	}

	parent[pRoot][0] = qRoot
	parent[qRoot][1] += parent[pRoot][1]
	if parent[qRoot][1] > *maxSize {
		*maxSize = parent[qRoot][1]
	}
}

func find(parent map[int][]int, p int) int {
	for p != parent[p][0] {
		parent[p] = parent[parent[p][0]]
		p = parent[p][0]
	}
	return p
}

func longestConsecutive(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	parent:= map[int][]int{}
	res := 1
	for _, e := range nums {
		if _, ok := parent[e]; ok {
			continue
		} else {
			parent[e] = []int{e, 1}
		}

		_, prevOk := parent[e-1]
		_, nextOk := parent[e+1]
		if prevOk {
			union(parent, e, e-1, &res)
		}

		if nextOk {
			union(parent, e, e+1, &res)
		}
	}
	return res
}
