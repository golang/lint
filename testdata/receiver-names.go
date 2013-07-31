// Test for bad receiver names.

// Package foo ...
package foo

type foo struct{}

func (this foo) f1() { // MATCH /should be a reflection of its identity/
}

func (self foo) f2() { // MATCH /should be a reflection of its identity/
}

func (f foo) f3() {
}

func (foo) f4() {
}
