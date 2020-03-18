package array

import "strconv"

// Array array struct
type Array struct {
	data []int
	size int
}

// CreateArray init array
func CreateArray(cap int) *Array {
	arr := new(Array)
	arr.data = make([]int, cap)
	arr.size = 0
	return arr
}

func (arr *Array) String() string {
	str := "cap=" + strconv.Itoa(cap(arr.data)) + ", size=" + strconv.Itoa(arr.size) + "\n["
	for i := 0; i < arr.size; i++ {
		str += strconv.Itoa(arr.data[i])
		if i != arr.size-1 {
			str += ", "
		}
	}
	return str + "]"
}
