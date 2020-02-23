package leetcode

/**
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.



Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
Note:

Although the above answer is in lexicographical order, your answer could be in any order you want.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func combineLetter(res *[]string, digits, letter string, i, end int, m *map[uint8][]string) {
	if i == end {
		*res = append(*res, letter)
		return
	}
	for _, e := range (*m)[digits[i]] {
		combineLetter(res, digits, letter+e, i+1, end, m)
	}
}

func letterCombinations(digits string) []string {
	m := map[uint8][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	res := []string{}
	combineLetter(&res, digits, "", 0, len(digits), &m)
	return res
}
