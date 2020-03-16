package math

import "strings"

/**
Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-binary
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func addBinary(a string, b string) string {
	m, n := len(a), len(b)
	if m < n {
		return addBinarySolution(&a, &b, m, n)
	}
	return addBinarySolution(&b, &a, n, m)
}

func addBinarySolution(s, l *string, sl, ll int) string {
	carry := 0
	temp := []string{"0", "1"}
	res := make([]string, ll+1)
	for i := 1; i < ll; i++ {
		res[i] = string((*l)[i-1])
	}
	for i, j := sl-1, ll-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if (*s)[i] == (*l)[j] {
			res[j+1] = temp[carry]
			if (*s)[i] == '1' {
				carry = 1
			} else {
				carry = 0
			}
		} else {
			res[j+1] = temp[1-carry]
		}
	}

	for i := ll - sl - 1; i >= 0 && carry == 1; i-- {
		if (*l)[i] == '1' {
			res[i+1] = "0"
		} else {
			res[i+1] = "1"
			carry = 0
		}
	}

	if carry == 1 {
		res[0] = "1"
	}
	return strings.Join(res, "")
}
