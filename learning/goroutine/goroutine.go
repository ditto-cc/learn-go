package main

import (
	"fmt"
	"time"
)

func main() {
	a := make([]int, 10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
