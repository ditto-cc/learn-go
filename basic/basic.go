package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 12
	bb = true
	ss = "gg"
)

// zero value
func varibleZeroVal() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

// assign value
func varibleInitVal() {
	var a int = 3
	var s string = "abc"
	fmt.Printf("%d %q\n", a, s)
}

// Type deduction
func varibleTypeDeduction() {
	var a, b, c, s = 1, 2, true, "abc"
	fmt.Println(a, b, c, s)
}

// Shorter Defination
func varibleConsice() {
	a, b, c, s := 3, 4, true, "abc"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	fmt.Printf("%.3f\n", cmplx.Pow(math.E, 1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const (
		filename string = "abc.txt"
		a, b            = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		java
		python
		golang
	)
	fmt.Println(cpp, java, python, golang)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello, World")
	varibleZeroVal()
	varibleInitVal()
	varibleTypeDeduction()
	varibleConsice()
	fmt.Println(aa, bb, ss)

	euler()
	triangle()
	consts()
	enums()
}
