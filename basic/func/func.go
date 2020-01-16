package main

import "fmt"

import "reflect"

import "runtime"

import "math"

func eval(a, b int, op string) (int, error) {
	res := 0
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res, _ = div(a, b)
	default:
		return 0, fmt.Errorf("unsupported operation: " + op)
	}
	return res, nil
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func sum(values ...int) int {
	s := 0
	for i := range values {
		s += values[i]
	}
	return s
}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	if res, err := eval(13, 3, "/"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	q, r := div(13, 3)
	fmt.Println(q, r)

	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5, 6))

	a, b := 3, 4
	b, a = swap(a, b)
	fmt.Println(a, b)
}
