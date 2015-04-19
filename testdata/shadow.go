// Test for shadowing variables

// Package foo ...
package foo

import "math"

func f(x int) bool {
	y := math.Sin(5)
	err := "test"
	math := "duck" // MATCH /shadowing variable - math/
	if x > 10 {
		x := 7 // MATCH /shadowing variable - x/
		y := 7 // MATCH /shadowing variable - y/
		if true {
			x = 7 // Fine
		}

		err := "notetst" // Fine since err is in ignore list
		return true
	}
	func(x int) { // MATCH /shadowing variable - x/
		return
	}(10)
	return false
}

var a int

func g(a int) int { // MATCH /shadowing variable - a/
	return a
}
func h() (a int) { // MATCH /shadowing variable - a/
	return a
}
func i(b int) (c int) {
	if true {
		b := 1 // MATCH /shadowing variable - b/
		c := 1 // MATCH /shadowing variable - c/
		_ = b
		_ = c
	}
	return
}
