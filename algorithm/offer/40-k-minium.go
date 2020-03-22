package offer

/**
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。



示例 1：

输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
示例 2：

输入：arr = [0,1,2,1], k = 1
输出：[0]


限制：

0 <= k <= arr.length <= 10000
0 <= arr[i] <= 10000
*/

func down(data []int, i, end int) {
	for i*2+1 < end {
		c := i*2 + 1
		if c+1 < end && data[c] < data[c+1] {
			c++
		}

		if data[i] >= data[c] {
			break
		}

		data[i], data[c] = data[c], data[i]
		i = c
	}
}

func getLeastNumbers(arr []int, k int) []int {
	data := arr[:k]
	if k == 0 {
		return data
	}

	for i := (k - 1) / 2; i >= 0; i-- {
		down(data, i, k)
	}

	n := len(arr)
	for i := k; i < n; i++ {
		if arr[i] < data[0] {
			data[0] = arr[i]
			down(data, 0, k)
		}
	}

	return data
}
