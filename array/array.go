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

/*
Filter the Array according to the result of the func
*/
func (array Array) Filter(callback func(val A, index int, array Array) bool) Array {
	newArray := Array{}
	for index, val := range array {
		if callback(val, index, array) {
			newArray = append(newArray, val)
		}
	}
	return newArray
}

/*
Map creates a new array with the results of calling a provided function on every element in this array.
*/
func (array Array) Map(callback func(val A, index int, array Array) A) Array {
	newArray := make(Array, len(array))
	for index, val := range array {
		newArray[index] = callback(val, index, array)
	}
	return newArray
}

/*
IndexOf return first index of the value
*/
func (array Array) IndexOf(value A) int {
	for i, v := range array {
		if value == v {
			return i
		}
	}
	return -1
}
