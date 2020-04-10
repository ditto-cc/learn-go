package main

import (
	"math/rand"
	"sort"
)

func IsSorted(arr []int) bool {
	return sort.IntsAreSorted(arr)
}

func GenerateRandomArray(ranL, ranR, n int) []int {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = rand.Intn(ranR-ranL) + ranL
	}
	return arr
}

func GenerateNearlyOrderedArray(n, swap int) []int {
	arr := make([]int, n)

	for i := range arr {
		arr[i] = i
	}

	for ; swap > 0; swap-- {
		i, j := rand.Intn(n), rand.Intn(n)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func CopyArray(data []int) []int {
	arr := make([]int, len(data))
	copy(arr, data)
	return arr
}
