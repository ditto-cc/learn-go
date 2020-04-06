package string

import (
	"sort"
	"strings"
)

/**
In English, we have a concept called root, which can be followed by some other words to form another longer word - let's call this word successor. For example, the root an, followed by other, which can form another word another.

Now, given a dictionary consisting of many roots and a sentence. You need to replace all the successor in the sentence with the root forming it. If a successor has many roots can form it, replace it with the root with the shortest length.

You need to output the sentence after the replacement.

Example 1:

Input: dict = ["cat", "bat", "rat"]
sentence = "the cattle was rattled by the battery"
Output: "the cat was rat by the bat"


Note:

The input will only have lower-case letters.
1 <= dict words number <= 1000
1 <= sentence words number <= 1000
1 <= root length <= 100
1 <= sentence words length <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/replace-words
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func replaceWords(dict []string, sentence string) string {
	m := map[string]bool{}
	counter := map[int]bool{}
	for _, str := range dict {
		m[str] = true
		counter[len(str)] = true
	}

	lengthSlice := make([]int, 0, len(counter))
	for k := range counter {
		lengthSlice = append(lengthSlice, k)
		delete(counter, k)
	}
	sort.Ints(lengthSlice)
	n := len(lengthSlice)

	strSlice := strings.Split(sentence, " ")
	for i, str := range strSlice {
		for j := 0; j < n; j++ {
			if lengthSlice[j] >= len(str) {
				break
			}
			if _, ok := m[str[:lengthSlice[j]]]; ok {
				strSlice[i] = str[:lengthSlice[j]]
				break
			}
		}
	}

	return strings.Join(strSlice, " ")
}
