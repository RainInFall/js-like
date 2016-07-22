package array2

// template type Array(A, B)
type A int
type B float32

//Array has js-like functions
type ArrayA []A
type ArrayB []B

/*
Map creates a new array with the results of calling a provided function on every element in this array.
*/
func (array ArrayA) Map(callback func(val A, index int, array ArrayA) B) ArrayB {
	newArray := make(ArrayB, len(array))
	for index, val := range array {
		newArray[index] = callback(val, index, array)
	}
	return newArray
}
