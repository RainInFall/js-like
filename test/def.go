package test

type Int struct {
	Value int
}

func (array ArraypInt) Len() int {
	return len(array)
}

func (array ArraypInt) Less(i, j int) bool {
	return false
}

func (array ArraypInt) Swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}
