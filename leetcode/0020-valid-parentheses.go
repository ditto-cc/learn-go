package leetcode

/**
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func isValid(s string) bool {
	stack := []rune{}
	for _, c := range s {
		if c == '{' || c == '[' || c == '(' {
			stack = append(stack, c)
		} else {
			l := len(stack)
			if l == 0 {
				return false
			}
			top := stack[l-1]
			stack = stack[:l-1]
			if (c == ')' && top != '(') ||
				(c == ']' && top != '[') ||
				(c == '}' && top != '{') {
				return false
			}
		}
	}
	return len(stack) == 0
}