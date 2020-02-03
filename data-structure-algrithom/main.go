package main

import (
	"fmt"
	"learn-go/data-structure-algrithom/data-structure/bst"
)

func visit(val int) {
	fmt.Printf("%d ", val)
}

func main() {

	arr := []int{4, 2, 3, 1, 5, 7}
	tree := bst.CreateBST()

	for _, e := range arr {
		tree.Add(e)
	}
	tree.Add(0)
	fmt.Println(tree)

	tree.InOrderNR(visit)
	fmt.Println()
	tree.PreOrderNR(visit)
	fmt.Println()
	tree.PostOrderNR(visit)
	fmt.Println()
	tree.LevelOrder(visit)
}
