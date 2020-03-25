package main

import (
	"fmt"
	"sync"
)

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d, Receive %c\n", id, n)
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func chanDemo() {
	//var c chan int // nil == c
	workers := make([]worker, 10)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)
	for i, w := range workers {
		w.in <- 'a' + i
	}

	for i, w := range workers {
		w.in <- 'A' + i
	}
	wg.Wait()
}

func main() {
	chanDemo()
}
