package string

/**
Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).

Example:

Input: S = "ADOBECODEBANC", T = "ABC"
Output: "BANC"
Note:

If there is no such window in S that covers all characters in T, return the empty string "".
If there is such window, you are guaranteed that there will always be only one unique minimum window in S.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-window-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func minWindow(s string, t string) string {
	res := s
	n, m := len(s), len(t)
	counterMap := map[byte]int{}
	windowMap := map[byte]int{}
	for i := 0; i < m; i++ {
		counterMap[t[i]]++
	}

	c := 0
	for l, r := 0, 0; r < n; r++ {
		if num, ok := counterMap[s[r]]; ok {
			windowMap[s[r]]++
			if windowMap[s[r]] <= num {
				c++
			}

			if c == m {
				for ; l < r; l++ {
					if windowMap[s[l]] == 0 {
						continue
					}
					if windowMap[s[l]] <= counterMap[s[l]] {
						break
					}
					windowMap[s[l]]--
				}

				if len(res) > r-l+1 {
					res = s[l : r+1]
				}
			}
		}
	}

	if c < m {
		return ""
	}
	return res
}
