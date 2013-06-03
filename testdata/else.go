// Test of return+else warning.

// Package pkg ...
package pkg

import "log"

func f(x int) bool {
	if x > 0 {
		return true
	} else { // MATCH /if.*return.*else.*outdent/
		log.Printf("non-positive x: %d", x)
	}
	return false
}
