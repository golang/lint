// Test for errors compared to strings.

// Package foo ...
package foo

import "strings"

func f(err error) bool {
	if strings.Contains(err.Error(), "a string") { // MATCH /should not be compared to string/
		return true
	}
}
