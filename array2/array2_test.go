package array2

import (
	"reflect"
	"testing"

	"github.com/RainInFall/assert"
)

func TestMap(t *testing.T) {
	assert.Init(t)

	func() {
		array := ArrayA{}
		newArray := array.Map(func(val A, index int, array ArrayA) B {
			assert.Ok(false)
			return 0
		})
		assert.Ok(reflect.DeepEqual(ArrayB{}, newArray))
	}()

	func() {
		array := ArrayA{1, 2, 3, 4, 5, 6}
		newArray := array.Map(func(val A, index int, array ArrayA) B {
			return B(val)
		})
		assert.Ok(reflect.DeepEqual(ArrayB{1, 2, 3, 4, 5, 6}, newArray))
	}()

	func() {
		array := ArrayA{1, 2, 3, 4, 5, 6}
		newArray := array.Map(func(val A, index int, array ArrayA) B {
			return B(val + 1)
		})
		assert.Ok(reflect.DeepEqual(newArray, ArrayB{2, 3, 4, 5, 6, 7}))
	}()
}
