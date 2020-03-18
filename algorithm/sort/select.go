package sort

func SelectionSort(data []int) {
	n := len(data) - 1
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i; j <= n; j++ {
			if data[j] < data[minIndex] {
				minIndex = j
			}
		}
		data[i], data[minIndex] = data[minIndex], data[i]
	}
}
