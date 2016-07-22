package object

import (
	"testing"

	"github.com/RainInFall/assert"
)

func (obj Object) Less(i, j First) bool {
	return obj[i] < obj[j]
}

func TestKeys(t *testing.T) {
	assert.Init(t)

	func() {
		obj := Object{
			0: "Zero",
			1: "RainInFall",
			3: "Leo.Xiongs",
			2: "Canyd",
		}

		keys := obj.Keys()

		assert.Ok(len(keys) == 4)
		occurs := make([]bool, 4, 4)
		for _, v := range keys {
			if occurs[v] {
				t.FailNow()
			}
			occurs[v] = true
		}
	}()
}
