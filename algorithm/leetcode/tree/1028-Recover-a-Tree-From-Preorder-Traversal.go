package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverFromPreorder(S string) *TreeNode {
	i, nd := 0, 0
	return recoverTree(&S, len(S), 0, &i, &nd)
}

func recoverTree(S *string, size, d int, i, nd *int) *TreeNode {
	if *i > size {
		return nil
	}

	val := 0
	for ; *i < size && (*S)[*i] != '-'; (*i)++ {
		val = val * 10 + int((*S)[*i] - '0')
	}

	node := &TreeNode{Val: val}

	j := *i
	for ; j < size && (*S)[j] == '-'; j++ {}
	depth := j - *i

	if d >= depth {
		*i = j
		*nd = depth
		return node
	}

	*i = j
	node.Left = recoverTree(S, size, depth, i, nd)
	if depth == *nd {
		node.Right = recoverTree(S, size, depth, i, nd)
	}

	return node
}