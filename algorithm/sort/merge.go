package sort

func merge(data []int, l, mid, r int) {
	size := r - l + 1
	temp := make([]int, size)

	i, j := l, mid+1
	for k := 0; k < size; k++ {
		if i > mid {
			temp[k] = data[j]
			j++
		} else if j > r {
			temp[k] = data[i]
			i++
		} else if data[i] < data[j] {
			temp[k] = data[i]
			i++
		} else {
			temp[k] = data[j]
			j++
		}
	}

	copy(data[l:r+1], temp)
}

func mergeSort(data []int, l, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2
	mergeSort(data, l, mid)
	mergeSort(data, mid+1, r)
	merge(data, l, mid, r)
}

func MergeSort(data []int) {
	mergeSort(data, 0, len(data)-1)
}
