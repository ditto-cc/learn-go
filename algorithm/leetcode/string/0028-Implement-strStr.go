package string

/**
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-strstr
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func strStr(haystack string, needle string) int {
	// KMP
	m, n := len(haystack), len(needle)
	next := make([]int, n+1)
	for i, j := 0, -1; i < n; {
		if j == -1 || needle[i] == needle[j] {
			i++
			j++
			next[j] = j
		} else {
			j = next[j]
		}
	}

	for i, j := 0, 0; i < m; {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}

		if j == n {
			return i - j
		}
	}
	return -1
}
