// Test for defer in loop.

// Package foo ...
package foo

func b() {
}

func f() {
	defer b()
	for {
		defer b() // MATCH /should not use defer in for loop/
	}
	defer b()
	for {
		defer b() // MATCH /should not use defer in for loop/
	}
	if true {
		defer b()
	}
}
