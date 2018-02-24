// Test that blank imports in package main are not flagged.
// OK

// Binary foo does something.
package main

import _ "fmt"

import (
	"os"
	_ "path"
)

var _ os.File // for "os"
