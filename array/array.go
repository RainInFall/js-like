package array

// template type Array(A)
type A int

//Array has js-like functions
type Array []A

/*
Some
*/
func (array Array) Some(f func(A, int, Array) bool) bool {
	for index, value := range array {
		if f(value, index, array) {
			return true
		}
	}
	return false
}
