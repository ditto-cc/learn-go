package array

import (
	"bytes"
	"fmt"
)

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
	res := bytes.NewBuffer([]byte{'['})
	res.WriteString(fmt.Sprintf("cap=%v, size=%v\n", cap(arr.data), arr.size))
	for i := 0; i < arr.size; i++ {
		res.WriteString(fmt.Sprintf("%v", arr.data[i]))
		if i != arr.size-1 {
			res.Write([]byte{',', ' '})
		}
	}
	res.WriteByte(']')
	return res.String()
}
