package string

import (
	"sort"
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
	m := map[string]int{}
	res := [][]string{}
	lastIndex := 0
	for _, str := range strs {
		strArr := strings.Split(str, "")
		sort.Strings(strArr)
		sortedStr := strings.Join(strArr, "")
		if index, ok := m[sortedStr]; ok {
			res[index] = append(res[index], str)
		} else {
			res = append(res, []string{str})
			m[sortedStr] = lastIndex
			lastIndex++
		}
	}

	return res
}
