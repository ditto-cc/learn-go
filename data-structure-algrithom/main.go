package main

import (
	"fmt"
	"learn-go/data-structure-algrithom/data-structure/array"
)

func visit(val int) {
	fmt.Printf("%d ", val)
}

func main() {
	// tree := bst.NewBSTree()
	// arr := []int{4, 2, 1, 3, 6, 5, 7}
	// for e := range arr {
	// 	tree.Add(e)
	// }
	// tree.InOrder(visit)
	// fmt.Println()
	// tree.PreOrder(visit)
	// fmt.Println()
	// tree.PostOrder(visit)
	// fmt.Println()
	arr := array.CreateArray(10)
	fmt.Println(arr)

	for i := 0; i < 5; i++ {
		arr.Append(i*2 + 2)
		fmt.Println(arr)
		arr.Prepend(i*2 + 1)
		fmt.Println(arr)
	}

	arr.PopLeft()
	fmt.Println(arr)
	arr.PopRight()
	fmt.Println(arr)
}
