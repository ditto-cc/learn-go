package backtracking

import "bytes"

/**
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	buf := bytes.NewBufferString("")
	gen(n, n, buf, &res)
	return res
}

func gen(l, r int, buf *bytes.Buffer, res *[]string) {
	if l == 0 {
		for ; r > 0; r-- {
			buf.WriteByte(')')
		}
		*res = append(*res, buf.String())
		return
	}

	buf.WriteByte('(')
	gen(l-1, r, buf, res)
	buf.Truncate(buf.Len() - 1)
	if l < r {
		buf.WriteByte(')')
		gen(l, r-1, buf, res)
		buf.Truncate(buf.Len() - 1)
	}
}
