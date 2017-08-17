// Test for deprecation warnings.

// Package main ...
package main

import "net/http/httputil"

func main() {
	httputil.NewClientConn(nil, nil) // MATCH /deprecated/
}
