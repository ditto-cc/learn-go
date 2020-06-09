package offer

import (
	"math"
)

/*
给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

 

示例 1:

输入: 12258
输出: 5
解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"
 

提示：

0 <= num < 231

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func translateNum(num int) int {
	if num < 10 {
		return 1
	}

	digitNum := int(math.Floor(math.Log10(float64(num)))) + 1
	digits, dp := make([]int, digitNum), make([]int, digitNum)

	for i := digitNum - 1; i >= 0; i, num = i-1, num/10 {
		digits[i] = num % 10
	}

	dp[0], dp[1] = 1, 1
	if digits[0] != 0 && digits[0] * 10 + digits[1] < 26 {
		dp[1] = 2
	}

	for i := 2; i < digitNum; i++ {
		dp[i] = dp[i-1]
		if digits[i-1] != 0 && digits[i-1] * 10 + digits[i] < 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[digitNum-1]
}
