package string

/**
Given an input string, reverse the string word by word.



Example 1:

Input: "the sky is blue"
Output: "blue is sky the"
Example 2:

Input: "  hello world!  "
Output: "world! hello"
Explanation: Your reversed string should not contain leading or trailing spaces.
Example 3:

Input: "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.


Note:

A word is defined as a sequence of non-space characters.
Input string may contain leading or trailing spaces. However, your reversed string should not contain leading or trailing spaces.
You need to reduce multiple spaces between two words to a single space in the reversed string.


Follow up:

For C programmers, try to solve it in-place in O(1) extra space.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-words-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

import (
	"strings"
)

func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	n, end := len(s), 1
	if n == 0 {
		return ""
	}

	res := []byte(s)
	for end = 1; end < n && res[end] != ' '; end++ {
	}

	for i := end + 1; i < n; {
		for ; i < n && res[i] == ' '; i++ {
		}
		for end++; i < n && res[i] != ' '; i, end = i+1, end+1 {
			res[end] = res[i]
		}
		if end < n {
			res[end] = ' '
		}
	}

	for i := 0; i < end/2; i++ {
		res[i], res[end-i-1] = res[end-i-1], res[i]
	}

	for i := 0; i < end; i++ {
		j := i
		for ; i < end && res[i] != ' '; i++ {
		}
		for l, r := j, i-1; l < r; l, r = l+1, r-1 {
			res[l], res[r] = res[r], res[l]
		}
	}

	return string(res[:end])

	//s = strings.Trim(s, " ")
	//if s == "" {
	//	return ""
	//}
	//split := strings.Split(s, " ")
	//buf := bytes.NewBuffer(nil)
	//for i := len(split) - 1; i >= 0; i-- {
	//	if split[i] == "" {
	//		continue
	//	}
	//	buf.WriteString(split[i])
	//	if i > 0 {
	//		buf.WriteByte(' ')
	//	}
	//}
	//return buf.String()
}
