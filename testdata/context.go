// Test occurrences of context.TODO() in a function that takes a context.Context
// argument.

// Package pkg does something.
package pkg

import (
	"context"
	context2 "context"

	contextFoo "example.com/contextFoo"

	context3 "golang.org/x/net/context"
)

func hasContext(ctx context.Context, b int, ctx2, ctx3 context.Context) context.Context {
	return context.TODO() // MATCH /should use one of \[ctx ctx2 ctx3\] instead of context.TODO\(\)/
}

func hasContext(ctx context2.Context, b int, ctx2, ctx3 context2.Context) context2.Context {
	return context2.TODO() // MATCH /should use one of \[ctx ctx2 ctx3\] instead of context2.TODO\(\)/
}

func hasContext3(ctx context3.Context, b int, ctx2, ctx3 context3.Context) context3.Context {
	return context3.TODO() // MATCH /should use one of \[ctx ctx2 ctx3\] instead of context3.TODO\(\)/
}

func hasContextFoo(ctx contextFoo.Context, b int, ctx2, ctx3 contextFoo.Context) contextFoo.Context {
	return contextFoo.TODO()
}

func hasFuncLitContext(ctx context2.Context) {
	func() {
		context2.TODO() // MATCH /should use one of \[ctx\] instead of context2.TODO\(\)/
	}()
}

func hasFuncLitContextArg() {
	func(ctx context2.Context) {
		context2.TODO() // MATCH /should use one of \[ctx\] instead of context2.TODO\(\)/
	}(context2.Background())
}
