package main

import (
	"testing"

	"github.com/RainInFall/assert"
)

//go:generate $GOPATH/bin/js-like array *Int

func TestComplie(t *testing.T) {
	assert.Init(t)

	array := ArraypInt{&Int{0}}
	array[0] = nil
	assert.Ok(len(array) == 1)
	assert.Ok(nil == array[0])
}
