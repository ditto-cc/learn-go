package string

/**
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-prefix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func longestCommonPrefix(strs []string) string {
	n := len(strs)

	if n == 0 {
		return ""
	} else if n == 1 {
		return strs[0]
	}

	prefix := strs[0]
	prefixLen := len(prefix)
	for k := 1; k < n && prefixLen > 0; k++ {
		strLen := len(strs[k])

		i, j := 0, 0
		for ; i < prefixLen && j < strLen && prefix[i] == strs[k][j]; i, j = i+1, j+1 {

		}

		prefixLen = i
	}

	return prefix[:prefixLen]
}
