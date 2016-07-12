package array

import (
	"testing"

	"github.com/RainInFall/assert"
)

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