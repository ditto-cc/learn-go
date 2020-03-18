package sort

func BubbleSort(data []int) {
	n, change := len(data), false
	for i := n; i > 0; i-- {
		change = false
		for j := 1; j < i; j++ {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
				change = true
			}
		}

		if !change {
			break
		}
	}
}
