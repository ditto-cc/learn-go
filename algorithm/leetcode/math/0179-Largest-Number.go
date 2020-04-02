package math

import (
	"fmt"
	"sort"
	"strings"
)

func largestNumber(nums []int) string {
	strSlice := make([]string, len(nums))
	for i, e := range nums {
		strSlice[i] = fmt.Sprintf("%d", e)
	}

	sort.Slice(strSlice, func(i, j int) bool {
		return strSlice[i]+strSlice[j] >= strSlice[j]+strSlice[i]
	})

	if strSlice[0] == "0" {
		return "0"
	}
	return strings.Join(strSlice, "")
}
