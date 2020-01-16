package array

func (arr *Array) Add(i, e int) {
	if arr.size >= len(arr.data) {
		arr.data = append(arr.data, 0)
	}
	for j := arr.size; j > i; j-- {
		arr.data[j] = arr.data[j-1]
	}
	arr.data[i] = e
	arr.size++
}

func (arr *Array) Append(e int) {
	arr.Add(arr.size, e)
}

func (arr *Array) Prepend(e int) {
	arr.Add(0, e)
}

func (arr *Array) Remove(i int) int {
	if i < 0 || i >= arr.size {
		panic("Illegal Index. Remove Failed.")
	}

	ret := arr.data[i]
	arr.data = append(arr.data[:i], arr.data[i+1:]...)
	arr.size--
	return ret
}

func (arr *Array) PopLeft() int {
	return arr.Remove(0)
}

func (arr *Array) PopRight() int {
	return arr.Remove(arr.size - 1)
}

func (arr *Array) Get(i int) int {
	if i < 0 || i >= arr.size {
		panic("Illegal Index. Get Failed.")
	}
	return arr.data[i]
}

func (arr *Array) Set(i, e int) {
	if i < 0 || i >= arr.size {
		panic("Illegal Index. Set Failed.")
	}
	arr.data[i] = e
}
