package array

/**
 * In a deck of cards, each card has an integer written on it.
 *
 * Return true if and only if you can choose X >= 2 such that it is possible to split the entire deck into 1 or more groups of cards, where:
 *
 * Each group has exactly X cards.
 * All the cards in each group have the same integer.
 *
 *
 * Example 1:
 *
 * Input: deck = [1,2,3,4,4,3,2,1]
 * Output: true
 * Explanation: Possible partition [1,1],[2,2],[3,3],[4,4].
 * Example 2:
 *
 * Input: deck = [1,1,1,2,2,2,3,3]
 * Output: false´
 * Explanation: No possible partition.
 * Example 3:
 *
 * Input: deck = [1]
 * Output: false
 * Explanation: No possible partition.
 * Example 4:
 *
 * Input: deck = [1,1]
 * Output: true
 * Explanation: Possible partition [1,1].
 * Example 5:
 *
 * Input: deck = [1,1,2,2,2,2]
 * Output: true
 * Explanation: Possible partition [1,1],[2,2],[2,2].
 *
 *
 * Constraints:
 *
 * 1 <= deck.length <= 10^4
 * 0 <= deck[i] < 10^4
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func hasGroupsSizeX(deck []int) bool {
	m := map[int]int{}
	for _, e := range deck {
		if _, ok := m[e]; ok {
			m[e]++
		} else {
			m[e] = 1
		}
	}
	res := 0
	for _, v := range m {
		if res == 0 {
			res = v
		} else if res >= 2 {
			res = gcd(res, v)
		} else {
			return false
		}
	}

	return res >= 2
}

func gcd(a, b int) int {
	if a == b {
		return a
	}

	r1, r2 := a&1, b&1
	switch {
	case r1 == 1 && r2 == 1:
		if a < b {
			return gcd(b-a, a)
		}
		return gcd(a-b, b)
	case r1 == 0 && r2 == 1:
		return gcd(a>>1, b)
	case r1 == 1 && r2 == 0:
		return gcd(a, b>>1)
	default:
		return gcd(a>>1, b>>1) << 1
	}
}
