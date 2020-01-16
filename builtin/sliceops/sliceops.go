package main

import "fmt"

func printSlice(arr []int) {
	fmt.Printf("arr=%v, len(arr)=%d, cap(arr)=%d\n", arr, len(arr), cap(arr))
}

func main() {
	fmt.Println("Creating Slice")
	var s []int
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, i*2+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	printSlice(s2)

	s3 := make([]int, 10, 32)
	printSlice(s3)

	fmt.Println("\nCopying Slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("\nDeleting Elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("\nPoping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println("front=", front)
	printSlice(s2)

	fmt.Println("\nPoping from back")
	back := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println("back=", back)
	printSlice(s2)

}
