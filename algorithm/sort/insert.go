package sort

func InsertionSort(data []int) {
	n := len(data)
	for i := 1; i < n; i++ {
		v := data[i]
		j := i
		for ; j > 0 && data[j-1] > v; j-- {
			data[j] = data[j-1]
		}
		data[j] = v
	}
}
