package sort

import "math/rand"

// quick sort 3 ways
func quickSort(data []int, l, r int) {
	if l >= r {
		return
	}

	index := rand.Intn(r-l+1) + l
	data[l], data[index] = data[index], data[l]

	lt, gt, i := l, r+1, l+1
	for i < gt {
		if data[i] == data[l] {
			i++
		} else if data[i] < data[l] {
			data[lt+1], data[i] = data[i], data[lt+1]
			i, lt = i+1, lt+1
		} else {
			data[gt-1], data[i] = data[i], data[gt-1]
			gt--
		}
	}
	data[l], data[lt] = data[lt], data[l]

	quickSort(data, l, lt-1)
	quickSort(data, gt, r)
}

func QuickSort(data []int) {
	quickSort(data, 0, len(data)-1)
}
