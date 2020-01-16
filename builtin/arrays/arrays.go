package main

import "fmt"

func printArray(arr []int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{3, 4, 5, 6, 7}
	var grid [4][5]bool

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for _, v := range arr3 {
		fmt.Println(v)
	}

	fmt.Println("Print Array(arr1)")
	printArray(arr1[:])
	fmt.Println(arr1)

	fmt.Println("Print Array(arr3)")
	printArray(arr3[:])
	fmt.Println(arr3)

}
