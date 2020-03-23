package utils

func Max(elements ...int) int {
	res := elements[0]
	for i := 1; i < len(elements); i++ {
		if elements[i] > res {
			res = elements[i]
		}
	}
	return res
}
