package main

import (
	"fmt"
	"learn-go/data-structure-algrithom/data-structure/bst"
	"learn-go/data-structure-algrithom/data-structure/entity"
	"learn-go/data-structure-algrithom/data-structure/queue"
	"learn-go/data-structure-algrithom/data-structure/stack"
)

func visit(s bst.Comparable) {
	fmt.Printf("%v ", s.(*entity.Student))
}

func testQueue(s queue.Queue) {
	for i := 0; i < 5; i++ {
		s.Enqueue(&entity.Student{Name: fmt.Sprintf("entity.Student%d", i), Score: float64(60.0 + float64(i)*0.2)})
		fmt.Println(s)
	}
}

func testStack(s stack.Stack) {
	for i := 0; i < 5; i++ {
		s.Push(&entity.Student{
			Name:  fmt.Sprintf("entity.Student%d", i),
			Score: float64(60.0 + float64(i)*0.2),
		})
		fmt.Println(s)
	}
}

func testBSTree(tree *bst.BSTree) {
	for i := 0; i < 5; i++ {
		tree.Add(&entity.Student{
			Name:  fmt.Sprintf("entity.Student%d", i),
			Score: float64(60.0 + float64(i)*0.2),
		})
		fmt.Println(tree)
	}
}

func main() {
	tree := bst.CreateBST()
	testBSTree(tree)
}
