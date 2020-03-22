package main

import (
	"bufio"
	"fmt"
	"learn-go/learning/functional/fib"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}

func WriteFib(filename string) {
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		if pathErr, ok := err.(*os.PathError); ok {
			fmt.Println(pathErr.Path, pathErr.Op, pathErr.Err)
		} else {
			fmt.Println("unknown error", err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fib()
	for i := 0; i < 30; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	WriteFib("fib.txt")
}
