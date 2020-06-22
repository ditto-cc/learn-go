package offer

/*
你有两个字符串，即pattern和value。 pattern字符串由字母"a"和"b"组成，用于描述字符串中的模式。例如，字符串"catcatgocatgo"匹配模式"aabab"（其中"cat"是"a"，"go"是"b"），该字符串也匹配像"a"、"ab"和"b"这样的模式。但需注意"a"和"b"不能同时表示相同的字符串。编写一个方法判断value字符串是否匹配pattern字符串。

示例 1：

输入： pattern = "abba", value = "dogcatcatdog"
输出： true
示例 2：

输入： pattern = "abba", value = "dogcatcatfish"
输出： false
示例 3：

输入： pattern = "aaaa", value = "dogcatcatdog"
输出： false
示例 4：

输入： pattern = "abba", value = "dogdogdogdog"
输出： true
解释： "a"="dogdog",b=""，反之也符合规则
提示：

0 <= len(pattern) <= 1000
0 <= len(value) <= 1000
你可以假设pattern只包含字母"a"和"b"，value仅包含小写字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pattern-matching-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func patternMatching(pattern string, value string) bool {
	if pattern == "" {
		return value == ""
	}

	vLen := len(value)
	pCount := []int{0, 0}
	aFirst, bFirst := -1, -1
	for i, c := range pattern {
		if c == 'a' {
			pCount[0]++
			if aFirst < 0 {
				aFirst = i
			}
		} else {
			pCount[1]++
			if bFirst < 0 {
				bFirst = i
			}
		}
	}

	if value == "" {
		return pCount[0] == 0 || pCount[1] == 0
	}

	if pCount[0] == 0 || pCount[1] == 0 {
		c := pCount[0]
		if c == 0 {
			c = pCount[1]
		}
		if vLen % c != 0 {
			return false
		}
		l := vLen / c
		s := value[:l]
		for i := l; i < vLen; {
			end := i + l
			if s != value[i:end] {
				return false
			}
			i = end
		}
		return true
	}

	var lenList [][]int
	for v, i := vLen, 0; v >= 0; i, v = i+1, vLen-pCount[1]*(i+1) {
		if v%pCount[0] != 0 {
			continue
		}
		lenList = append(lenList, []int{v / pCount[0], i})
	}

	for _, e := range lenList {
		aLen, bLen := e[0], e[1]
		var aStr, bStr string
		if aFirst == 0 {
			aStr = value[:aLen]
			index := bFirst * aLen
			bStr = value[index : index+bLen]
		} else {
			bStr = value[:bLen]
			index := aFirst * bLen
			aStr = value[index : index+aLen]
		}

		i, end := 0, 0
		for _, c := range pattern {
			if c == 'a' {
				end = i+aLen
				if aStr != value[i:end] {
					break
				}
				i = end
			} else {
				end = i+bLen
				if bStr != value[i:end] {
					break
				}
				i = end
			}
		}
		if i == vLen {
			return true
		}
	}

	return false
}