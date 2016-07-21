package array

import (
	"sort"
)

// template type Array(A)
type A int

//Array has js-like functions
type Array []A

/*
Some tests whether some element in the array passes the test implemented by the provided function
*/
func (array Array) Some(f func(A, int, Array) bool) bool {
	for index, value := range array {
		if f(value, index, array) {
			return true
		}
	}
	return false
}

/*
Reverse the Array
*/
func (array Array) Reverse() Array {
	mid := len(array) / 2
	for i, j := 0, len(array)-1; i < mid; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}

/*
Sort the Array
*/
func (array Array) Sort() Array {
	sort.Sort(array)
	return array
}
