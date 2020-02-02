package main

import (
	"fmt"
	"learn-go/data-structure-algrithom/data-structure/queue"
)

func visit(val int) {
	fmt.Printf("%d ", val)
}

func main() {

	q1 := queue.CreateArrayQueue()
	q2 := queue.CreateListQueue()

	for i := 0; i < 10; i++ {
		q1.Enqueue(i)
		q2.Enqueue(i)
		fmt.Println(q1)
		fmt.Println(q2)
	}

}
