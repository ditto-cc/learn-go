package leetcode

import (
	"strings"
)

/**
Given an array of strings, group anagrams together.

Example:

Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
Output:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
Note:

All inputs will be in lowercase.
The order of your output does not matter.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/group-anagrams
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func groupAnagrams(strs []string) [][]string {
	chars := make([]string, 26, 26)
	m := map[string][]string{}
	res := [][]string{}
	for _, str := range strs {
		for i := range chars {
			chars[i] = ""
		}
		for _, c := range str {
			index := int(c) - 97
			chars[index] += string(c)
		}
		key := strings.Join(chars, "")
		m[key] = append(m[key], str)
	}
	for _, value := range m {
		res = append(res, value)
	}
	return res
}
