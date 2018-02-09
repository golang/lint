// Test that exported names have correct comments.

// Package pkg does something.
package pkg

import "time"

type T int // MATCH /exported type T.*should.*comment.*or.*unexport/

func (T) F() {} // MATCH /exported method T\.F.*should.*comment.*or.*unexport/

// this is a nice type.
// MATCH /comment.*exported type U.*should.*form.*"U ..."/
type U string

// EasyJSON struct is used for json Marshal/Unmarshal.
//easyjson:json
type EasyJSON struct {
	U
}

// this is a neat function.
// MATCH /comment.*exported method U\.G.*should.*form.*"G ..."/
func (U) G() {}

// G is a neat function. The comment ends in a period.
func (U) G() {}

// G is a neat function. The comment ends in a question mark?
func (U) G() {}

// A V is a string.
type V string

// VVar is a string.
var VVar V

// V.H has a pointer receiver

func (*V) H() {} // MATCH /exported method V\.H.*should.*comment.*or.*unexport/

var W = "foo" // MATCH /exported var W.*should.*comment.*or.*unexport/

const X = "bar" // MATCH /exported const X.*should.*comment.*or.*unexport/

var Y, Z int // MATCH /exported var Z.*own declaration/

// Location should be okay, since the other var name is an underscore.
var Location, _ = time.LoadLocation("Europe/Istanbul") // not Constantinople

// this is improperly documented
// MATCH /comment.*const.*Thing.*form.*"Thing ..."/
const Thing = "wonderful"
