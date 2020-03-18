package sort

func HeapSort(data []int) {
	n := len(data)

	for i := (n - 1) / 2; i >= 0; i-- {
		down(data, i, n)
	}

	for i := 1; i < n; i++ {
		data[n-i], data[0] = data[0], data[n-i]
		down(data, 0, n-i)
	}
}

func down(data []int, i, end int) {
	for i*2+1 < end {
		c := i*2 + 1
		if c+1 < end && data[c+1] > data[c] {
			c++
		}

		if data[i] >= data[c] {
			break
		}

		data[i], data[c] = data[c], data[i]
		i = c
	}
}
