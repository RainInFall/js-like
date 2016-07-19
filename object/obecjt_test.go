package object

import (
	"testing"

	"github.com/RainInFall/assert"
)

func TestKeys(t *testing.T) {
  assert.Init(t)

  func(){
    obj := Object{
      1 : "RainInFall",
      3 : "Leo.Xiongs",
      2 : "Canyd",
    }

    Keys(obj)
  }()
}
