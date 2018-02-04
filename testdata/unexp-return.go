// Test for unexported return types.

// Package foo ...
package foo

type hidden struct{}

// Returns a hidden type, which is annoying.
func Exported() hidden { // MATCH /Exported.*unexported.*hidden/
	return hidden{}
}

// ExpErr returns a builtin type.
func ExpErr() error { // ok
}

func (hidden) ExpOnHidden() hidden { // ok
}

// Description ..
type T struct{}

// Returns a hidden type, which is annoying.
func (T) MethodOnT() hidden { // MATCH /method MethodOnT.*unexported.*hidden/
	return hidden{}
}

// Returns a T.
func ExpT() T { // ok
	return T{}
}

func unexp() hidden { // ok
	return hidden{}
}

// This is slightly sneaky: we shadow the builtin "int" type.

type int struct{}

// Returns an unexported type from this package.
func ExportedIntReturner() int { // MATCH /ExportedIntReturner.*unexported.*int/
	return int{}
}
