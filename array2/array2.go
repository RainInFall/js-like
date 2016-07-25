package array2

// template type ArrayD(A, B, ArrayA, ArrayB)
type A int
type B float32
type ArrayA []A
type ArrayB []B

//Dummy
type ArrayD interface{}

/*
Map creates a new array with the results of calling a provided function on every element in this array.
*/
func (array ArrayA) Map_D_(callback func(val A, index int, array ArrayA) B) ArrayB {
	newArray := make(ArrayB, len(array))
	for index, val := range array {
		newArray[index] = callback(val, index, array)
	}
	return newArray
}
