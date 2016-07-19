package array

import (
	"sort"
	"testing"

	"github.com/RainInFall/assert"
)

func TestSort(t *testing.T) {
	assert.Init(t)

	func() {
		array := make(Array, 0, 10)

		assert.Ok(sort.IsSorted(array.Sort()))
	}()

	func() {
		array := Array{1, 5, 9, 8, 7, 6, 2, 4, 3}

		assert.Ok(sort.IsSorted(array.Sort()))
	}()
}

func TestEquals(t *testing.T) {
	assert.Init(t)

	func() {
		array0 := Array{}
		array1 := Array{}

		assert.Ok(array0.Equals(array1))
		assert.Ok(array1.Equals(array0))
	}()

	func() {
		array0 := Array{0, 2, 4}
		array1 := Array{0, 2, 4}

		assert.Ok(array0.Equals(array1))
		assert.Ok(array1.Equals(array0))

		array0[0] = 1
		assert.Ok(!array0.Equals(array1))
		assert.Ok(!array1.Equals(array0))
	}()

	func() {
		array0 := Array{0, 2, 4}
		array1 := Array{0, 2}

		assert.Ok(!array0.Equals(array1))
		assert.Ok(!array1.Equals(array0))

		array1 = append(array1, 4)
		assert.Ok(array0.Equals(array1))
		assert.Ok(array1.Equals(array0))
	}()
}

func TestRevese(t *testing.T) {
	assert.Init(t)

	func() {
		array := Array{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		assert.Ok(array.Reverse().Equals(Array{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}))
	}()
}

func TestSome(t *testing.T) {
	assert.Init(t)
	//Test empty array
	func() {
		array := make(Array, 0, 10)

		assert.Ok(!array.Some(func(value A, index int, _array Array) bool {
			t.Errorf("Should not be here")
			return false
		}))
	}()
	//Test return true
	func() {
		array := make(Array, 10, 10)

		assert.Ok(array.Some(func(value A, index int, _array Array) bool {
			return true
		}))
	}()
	//Test return false
	func() {
		array := make(Array, 10, 10)

		assert.Ok(!array.Some(func(value A, index int, _array Array) bool {
			return false
		}))
	}()
	//Test array is same from callback
	func() {
		array := make(Array, 10, 10)
		for i := range array {
			array[i] = A(7 - i)
		}
		array.Some(func(value A, index int, _array Array) bool {
			assert.Ok(array[index] == _array[index])
			_array[index] = A(0)
			return false
		})
		assert.Ok(!array.Some(func(value A, index int, _array Array) bool {
			return 0 != value
		}))
		for _, value := range array {
			assert.Ok(0 == value)
		}
	}()
}
