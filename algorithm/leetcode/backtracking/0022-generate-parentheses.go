package backtracking

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

func genParenthesis(n, c int, s string, res *[]string) {
	if n == 0 && c == 0 {
		*res = append(*res, s)
		return
	}

	if n != 0 {
		genParenthesis(n-1, c+1, s+"(", res)
	}
	if c != 0 {
		genParenthesis(n, c-1, s+")", res)
	}
}

func generateParenthesis(n int) []string {
	res := []string{}
	genParenthesis(n, 0, "", &res)
	return res
}
