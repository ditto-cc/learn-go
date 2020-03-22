package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(convert("LEETCODEISHIRING", 4))
}
func convert(s string, numRows int) string {
	n := len(s)
	if numRows <= 1 || n <= 2 {
		return s
	}

	pace, res := (numRows-1)<<1, bytes.Buffer{}
	for i := 0; i < numRows; i++ {
		for j := i; j < n; j += pace {
			res.WriteByte(s[j])

			index := j + pace - i*2
			if i > 0 && index < n && j != index {
				res.WriteByte(s[index])
			}
		}
	}

	return res.String()
}
