// Test for name linting.

// Package pkg_with_underscores ...
package pkg_with_underscores // MATCH /underscore.*package name/

var var_name int // MATCH /underscore.*var.*var_name/

type t_wow struct { // MATCH /underscore.*type.*t_wow/
	x_damn int      // MATCH /underscore.*field.*x_damn/
	Url    *url.URL // MATCH /struct field.*Url.*URL/
}

const fooId = "blah" // MATCH /fooId.*fooID/

func f_it() { // MATCH /underscore.*func.*f_it/
	more_underscore := 4 // MATCH /underscore.*var.*more_underscore/

	x := foo_proto.Blah{} // should be okay
}

// Common styles in other languages that don't belong in Go.
const (
	CPP_CONST   = 1 // MATCH /ALL_CAPS.*CamelCase/
	kLeadingKay = 2 // MATCH /k.*leadingKay/

	HTML  = 3 // okay; no underscore
	X509B = 4 // ditto
)
