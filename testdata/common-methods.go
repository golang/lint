// Test that we don't nag for comments on common methods.
// OK

// Package pkg ...
package pkg

import (
  "net/http"
)
// rename test
type Test int

func (Test) Error() string                                    { return "" }
func (Test) String() string                                   { return "" }
func (Test) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
func (Test) Read(p []byte) (n int, err error)                 { return 0, nil }
func (Test) Write(p []byte) (n int, err error)                { return 0, nil }
