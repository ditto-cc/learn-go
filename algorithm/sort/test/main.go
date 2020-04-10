package main

import (
	"fmt"
	"learn-go/algorithm/sort"
	"time"
)

func testSort(sortName string, data []int, sort func([]int)) {
	start := time.Now()
	sort(data)
	elapse := time.Since(start)
	fmt.Printf("%s Test, N = %d, Time = %v, Sorted = %v\n", sortName, len(data), elapse, IsSorted(data))
}

func main() {
	n := 100000
	//arr1 := GenerateRandomArray(0, n, n)
	arr1 := GenerateNearlyOrderedArray(n, 10)
	arr2 := CopyArray(arr1)
	arr3 := CopyArray(arr1)

	testSort("Insertion Sort", arr1, sort.InsertionSort)
	testSort("Bubble Sort", arr2, sort.BubbleSort)
	testSort("Selection Sort", arr3, sort.SelectionSort)

	n = 1000000
	arr1 = GenerateNearlyOrderedArray(n, 10)
	arr2 = CopyArray(arr1)
	arr3 = CopyArray(arr1)
	testSort("Quick Sort", arr1, sort.QuickSort)
	testSort("Heap Sort", arr2, sort.HeapSort)
	testSort("Merge Sort", arr3, sort.MergeSort)
}
