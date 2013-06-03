// Test that exported names have correct comments.

// Package pkg does something.
package pkg

type T int // MATCH /exported type T.*should.*comment/

func (T) F() {} // MATCH /exported method T\.F.*should.*comment/

// this is a nice type.
// MATCH /comment.*exported type U.*should.*form.*"U ..."/
type U string

// this is a neat function.
// MATCH /comment.*exported method U\.G.*should.*form.*"G ..."/
func (U) G() {}

// A V is a string.
type V string

// V.H has a pointer receiver

func (*V) H() {} // MATCH /exported method V\.H.*should.*comment/

var W = "foo" // MATCH /exported var W.*should.*comment/

const X = "bar" // MATCH /exported const X.*should.*comment/

var Y, Z int // MATCH /exported var Y.*own declaration/
