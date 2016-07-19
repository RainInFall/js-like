package array

import (
	"sort"
)

// template type Array(A)
type A int

//Array has js-like functions
type Array []A

/********Self defination Begin*******************/
func (array Array) Len() int {
	return len(array)
}

func (array Array) Less(i, j int) bool {
	return array[i] < array[j]
}

func (array Array) Swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}

/********Self defination End*******************/

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

/*
Equals compare two array with !=
*/
func (array Array) Equals(other Array) bool {
	if len(other) != len(array) {
		return false
	}
	for i, v := range array {
		if v != other[i] {
			return false
		}
	}
	return true
}
