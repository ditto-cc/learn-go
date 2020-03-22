package string

/**
Given a string which consists of lowercase or uppercase letters, find the length of the longest palindromes that can be built with those letters.

This is case sensitive, for example "Aa" is not considered a palindrome here.

Note:
Assume the length of given string will not exceed 1,010.

Example:

Input:
"abccccdd"

Output:
7

Explanation:
One longest palindrome that can be built is "dccaccd", whose length is 7.
*/

func longestPalindrome(s string) int {
	letter := make([]int, 52)
	for _, c := range s {
		switch {
		case 65 <= c && c <= 90:
			letter[c-65]++
		case 97 <= c && c <= 122:
			letter[c-71]++
		}
	}

	hasOdd, res := 0, 0
	for _, n := range letter {
		res += n - n&1
		hasOdd |= n & 1
	}

	return res + hasOdd
}
