package string

/**
Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false
Note:
You may assume the string contains only lowercase alphabets.

Follow up:
What if the inputs contain unicode characters? How would you adapt your solution to such case?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-anagram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isAnagram(s string, t string) bool {
	chars := make([]int, 26, 26)
	if len(s) != len(t) {
		return false
	}

	for _, c := range s {
		chars[int(c)-97] += 1
	}

	for _, c := range t {
		index := int(c) - 97
		if chars[index] == 0 {
			return false
		}
		chars[index] -= 1
	}
	for _, e := range chars {
		if e != 0 {
			return false
		}
	}
	return true

	// Followed Up
	//if len(s) != len(t) {
	//	return false
	//}
	//m := map[rune]int{}
	//for _, c := range s {
	//	if _, ok := m[c]; ok {
	//		m[c]++
	//	} else {
	//		m[c] = 1
	//	}
	//}
	//
	//for _, c := range t {
	//	if val, ok := m[c]; ok && val != 0 {
	//		m[c]--
	//	} else {
	//		return false
	//	}
	//}
	//
	//for _, val := range m {
	//	if val != 0 {
	//		return false
	//	}
	//}
	//return true
}
