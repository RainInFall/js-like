package main

import (
	"sort"
)

// template type Array(A)

//Array has js-like functions
type ArraypInt []*Int

/*
Some tests whether some element in the array passes the test implemented by the provided function
*/
func (array ArraypInt) Some(f func(*Int, int, ArraypInt) bool) bool {
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
func (array ArraypInt) Reverse() ArraypInt {
	mid := len(array) / 2
	for i, j := 0, len(array)-1; i < mid; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}

/*
Sort the Array
*/
func (array ArraypInt) Sort() ArraypInt {
	sort.Sort(array)
	return array
}
