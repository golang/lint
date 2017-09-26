// Test for not using fmt.Errorf or testing.Errorf.

// Package foo ...
package foo

import (
	"errors"
	"fmt"
	"testing"
)

func f(x int) error {
	if x > 10 {
		// MATCH /should replace.*errors\.New\(fmt\.Sprintf\(\.\.\.\)\).*fmt\.Errorf\(\.\.\.\)/ -> `		return fmt.Errorf("something %d", x)`
		return errors.New(fmt.Sprintf("something %d", x)) 
		
	}
	if x > 5 {
		return errors.New(g("blah")) // ok
	}
	if x > 4 {
		return errors.New("something else") // ok
	}
	return nil
}

// TestF is a dummy test
func TestF(t *testing.T) error {
	var x int = 1
	if x > 10 {
		return t.Error(fmt.Sprintf("something %d", x)) // MATCH /should replace.*t\.Error\(fmt\.Sprintf\(\.\.\.\)\).*t\.Errorf\(\.\.\.\)/
	}
	if x > 5 {
		return t.Error(g("blah")) // ok
	}
	if x > 4 {
		return t.Error("something else") // ok
	}
	return nil
}

func g(s string) string { return "prefix: " + s }
