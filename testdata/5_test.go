// This file ends in _test.go, so we should not warn about doc comments.
// OK

package pkg

import "testing"

type H int

func TestSomething(t *testing.T) {
}

func TestSomething_suffix(t *testing.T) {
}

type Buffer struct{}

func (Buffer) read() {}

func ExampleBuffer_read() {}
