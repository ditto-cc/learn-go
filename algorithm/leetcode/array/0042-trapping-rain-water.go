package array

/**
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.


The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped. Thanks Marcos for contributing this image!

Example:

Input: [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/trapping-rain-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func trap(height []int) int {
	n, res := len(height), 0

	if n < 3 {
		return res
	}

	for lMax, rMax, i, j := 0, 0, 0, n-1; i <= j; {
		if height[i] < height[j] {
			if height[i] < lMax {
				res += lMax - height[i]
			} else {
				lMax = height[i]
			}
			i++
		} else {
			if height[j] < rMax {
				res += rMax - height[j]
			} else {
				rMax = height[j]
			}
			j--
		}
	}

	return res
}
