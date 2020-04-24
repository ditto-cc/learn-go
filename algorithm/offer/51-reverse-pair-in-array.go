package offer

/*
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。



示例 1:

输入: [7,5,6,4]
输出: 5


限制：

0 <= 数组长度 <= 50000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func reversePairs(nums []int) int {
	n := len(nums)
	temp := make([]int, n)
	return mergeSort(nums, temp, 0, n)
}

func mergeSort(nums, temp []int, l, r int) int {
	if l >= r-1 {
		return 0
	}
	mid := int(uint(l+r) >> 1)
	res := 0
	res += mergeSort(nums, temp, l, mid)
	res += mergeSort(nums, temp, mid, r)
	return res + merge(nums, temp, l, mid, r)
}

func merge(nums, temp []int, l, mid, r int) int {
	res := 0

	for k, i, j := l, l, mid; k < r; k++ {
		if j >= r {
			temp[k] = nums[i]
			i++
		} else if i >= mid || nums[i] <= nums[j] {
			temp[k] = nums[j]
			j++
		} else {
			res += r - j
			temp[k] = nums[i]
			i++
		}
	}

	copy(nums[l:r], temp[l:r])

	return res
}
