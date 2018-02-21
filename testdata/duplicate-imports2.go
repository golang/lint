// Test that duplicate imports in library packages are flagged in an import section that includes blank imports.

// Package main...
package main

import (
  "strings"

  // Don't gripe about these next two lines.
  _ "compress/zlib"
  _ "syscall"

  /* MATCH /duplicate import: "things/cats" is imported 3 times/ */ aaa "things/cats"
  /* MATCH /duplicate import: "things/cats" is imported 3 times/ */ bbb "things/cats"
  /* MATCH /duplicate import: "things/cats" is imported 3 times/ */ ccc "things/cats"
)
