package string

/**
For strings S and T, we say "T divides S" if and only if S = T + ... + T  (T concatenated with itself 1 or more times)

Return the largest string X such that X divides str1 and X divides str2.



Example 1:

Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"
Example 2:

Input: str1 = "ABABAB", str2 = "ABAB"
Output: "AB"
Example 3:

Input: str1 = "LEET", str2 = "CODE"
Output: ""


Note:

1 <= str1.length <= 1000
1 <= str2.length <= 1000
str1[i] and str2[i] are English uppercase letters.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/greatest-common-divisor-of-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func gcdOfStrings(str1 string, str2 string) string {
	m, n := len(str1), len(str2)
	if m > n {
		str1, str2 = str2, str1
		m, n = n, m
	}

	for str1 != "" {
		for i := range str1 {
			if str1[i] != str2[i] {
				return ""
			}
		}
		str2 = str2[m:]
		if n-m < m {
			str2, str1 = str1, str2
			m, n = n-m, m
		} else {
			n = n - m
		}
	}

	return str2
}
