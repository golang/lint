// Copyright (c) 2013 The Go Authors. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or at
// https://developers.google.com/open-source/licenses/bsd.

// golint lints the Go source files named on its command line.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/lint"
)

var minConfidence = flag.Float64("min_confidence", 0.8, "minimum confidence of a problem to print it")

func main() {
	flag.Parse()

	// TODO(dsymonds): Support linting of stdin.
	for _, filename := range flag.Args() {
		if isDir(filename) {
			lintDir(filename)
		} else {
			lintFile(filename)
		}
	}
}

func isDir(filename string) bool {
	fi, err := os.Stat(filename)
	return err == nil && fi.IsDir()
}

func lintFile(filename string) {
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Failed reading %v: %v", filename, err)
		return
	}

	l := new(lint.Linter)
	ps, err := l.Lint(filename, src)
	if err != nil {
		log.Printf("Failed parsing %v: %v", filename, err)
		return
	}
	for _, p := range ps {
		if p.Confidence >= *minConfidence {
			fmt.Printf("%s:%v: %s\n", filename, p.Position, p.Text)
		}
	}
}

func lintDir(dirname string) {
	filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && strings.HasSuffix(path, ".go") {
			lintFile(path)
		}
		return err
	})
}
