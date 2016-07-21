package main

import (
	"testing"

	"github.com/RainInFall/assert"
)

//go:generate js-like array *Int

type Int struct {
	Value int
}

func (array ArraypInt) Len() int {
	return len(array)
}

func (array ArraypInt) Less(i, j int) bool {
	return false
}

func (array ArraypInt) Swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}

func TestComplie(t *testing.T) {
	assert.Init(t)

	array := ArraypInt{&Int{0}}
	array[0] = nil
	assert.Ok(len(array) == 1)
	assert.Ok(nil == array[0])
}
