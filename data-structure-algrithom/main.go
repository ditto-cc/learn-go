package main

import "fmt"

import "learn-go/data-structure-algrithom/data-structure/bst"

func visit(val int) {
	fmt.Printf("%d ", val)
}

func main() {
	tree := bst.NewBSTree()
	arr := []int{4, 2, 1, 3, 6, 5, 7}
	for e := range arr {
		tree.Add(e)
	}
	tree.InOrder(visit)
	fmt.Println()
	tree.PreOrder(visit)
	fmt.Println()
	tree.PostOrder(visit)
	fmt.Println()
}
