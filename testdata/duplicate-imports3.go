// Test that duplicate imports in library packages are flagged in an average-looking import section.

// Package main...
package main

import (
  "strings"
  /* MATCH /duplicate import: "stuff" is imported 2 times/ */  "stuff"

  // Don't gripe about these next two lines.
  _ "compress/zlib"
  _ "syscall"

  /* MATCH /duplicate import: "things/cats" is imported 2 times/ */ aaa "things/cats"
  /* MATCH /duplicate import: "things/cats" is imported 2 times/ */ bbb "things/cats"
)

/* MATCH /duplicate import: "stuff" is imported 2 times/ */ import "stuff"
