package array

import (
	"reflect"
	"sort"
	"testing"

	"github.com/RainInFall/assert"
)

func (array Array) Len() int {
	return len(array)
}

func (array Array) Less(i, j int) bool {
	return false
}

func (array Array) Swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}

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

func TestRevese(t *testing.T) {
	assert.Init(t)

	func() {
		array := Array{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		assert.Ok(reflect.DeepEqual(array.Reverse(), Array{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}))
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

func TestFilter(t *testing.T) {
	assert.Init(t)

	func() {
		array := Array{}
		newArray := array.Filter(func(val A, index int, array Array) bool {
			t.Fail()
			return false
		})
		assert.Ok(nil != newArray)
		assert.Ok(newArray.Len() == 0 && len(newArray) == 0)
		assert.Ok(reflect.TypeOf(array) == reflect.TypeOf(newArray))
	}()

	func() {
		array := Array{1, 2, 3, 4, 5, 6}
		newArray := array.Filter(func(val A, index int, array Array) bool {
			return false
		})
		assert.Ok(nil != newArray)
		assert.Ok(newArray.Len() == 0 && len(newArray) == 0)
		assert.Ok(reflect.TypeOf(array) == reflect.TypeOf(newArray))
	}()

	func() {
		array := Array{1, 2, 3, 4, 5, 6}
		newArray := array.Filter(func(val A, index int, array Array) bool {
			return 0 == val%2
		})

		assert.Ok(reflect.DeepEqual(newArray, Array{2, 4, 6}))
	}()
}

func TestMap(t *testing.T) {
	assert.Init(t)

	func() {
		array := Array{}
		newArray := array.Map(func(val A, index int, array Array) A {
			assert.Ok(false)
			return val
		})
		assert.Ok(reflect.DeepEqual(array, newArray))
	}()

	func() {
		array := Array{1, 2, 3, 4, 5, 6}
		newArray := array.Map(func(val A, index int, array Array) A {
			return val
		})
		assert.Ok(reflect.DeepEqual(array, newArray))
	}()

	func() {
		array := Array{1, 2, 3, 4, 5, 6}
		newArray := array.Map(func(val A, index int, array Array) A {
			return val * val
		})
		assert.Ok(reflect.DeepEqual(newArray, Array{1, 4, 9, 16, 25, 36}))
	}()
}
