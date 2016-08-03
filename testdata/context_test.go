// Test occurrences of context.TODO() in a function that takes a context.Context
// argument.

// Package pkg_test does something.
package pkg_test

import (
	"context"
	context2 "context"

	contextFoo "example.com/contextFoo"

	context3 "golang.org/x/net/context"
)

func TestBar() {
	context.TODO()  // MATCH /should use context.Background\(\) instead of context.TODO\(\) in tests/
	context2.TODO() // MATCH /should use context.Background\(\) instead of context2.TODO\(\) in tests/
	context3.TODO() // MATCH /should use context.Background\(\) instead of context3.TODO\(\) in tests/

	contextFoo.TODO()
}
