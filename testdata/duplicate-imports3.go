// Test that duplicate imports in library packages are flagged in an import section that includes blanks and a floating import statement.

// Package main...
package main

import (
  "strings"
  /* MATCH /duplicate import: "stuff" is imported 2 times/ */  "stuff"

  // Don't gripe about these next three lines.
  _ "compress/zlib"
  _ "syscall"
  _ "os"

  /* MATCH /duplicate import: "dogs" is imported 2 times/ */ _ "dogs"

  /* MATCH /duplicate import: "things/cats" is imported 2 times/ */ aaa "things/cats"
  /* MATCH /duplicate import: "things/cats" is imported 2 times/ */ bbb "things/cats"
)

/* MATCH /duplicate import: "stuff" is imported 2 times/ */ import "stuff"
  /* MATCH /duplicate import: "dogs" is imported 2 times/ */ import _ "dogs"
