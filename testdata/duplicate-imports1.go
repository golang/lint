// Test that duplicate imports in library packages are flagged in an average-looking import section.

// Package main...
package main

import (
  "strings"

  /* MATCH /duplicate import: "things/cats" is imported 2 times/ */ aaa "things/cats"
  /* MATCH /duplicate import: "things/cats" is imported 2 times/ */ bbb "things/cats"
)
