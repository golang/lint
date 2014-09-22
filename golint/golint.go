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
	"go/build"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/golang/lint"
)

var minConfidence = flag.Float64("min_confidence", 0.8, "minimum confidence of a problem to print it")

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\tgolint [flags] # runs on package in current directory\n")
	fmt.Fprintf(os.Stderr, "\tgolint [flags] package\n")
	fmt.Fprintf(os.Stderr, "\tgolint [flags] directory\n")
	fmt.Fprintf(os.Stderr, "\tgolint [flags] files... # must be a single package\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	switch flag.NArg() {
	case 0:
		os.Exit(lintDir("."))
	case 1:
		arg := flag.Arg(0)
		if isDir(arg) {
			os.Exit(lintDir(arg))
		} else if exists(arg) {
			os.Exit(lintFiles(arg))
		} else {
			os.Exit(lintPackage(arg))
		}
	default:
		os.Exit(lintFiles(flag.Args()...))
	}
}

func isDir(filename string) bool {
	fi, err := os.Stat(filename)
	return err == nil && fi.IsDir()
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func lintFiles(filenames ...string) int {
	files := make(map[string][]byte)
	for _, filename := range filenames {
		src, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		files[filename] = src
	}

	l := new(lint.Linter)
	ps, err := l.LintFiles(files)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return 1
	}
	result := 0
	for _, p := range ps {
		if p.Confidence >= *minConfidence {
			fmt.Printf("%v: %s\n", p.Position, p.Text)
			result = 1
		}
	}
	return result
}

func lintDir(dirname string) int {
	pkg, err := build.ImportDir(dirname, 0)
	return lintImportedPackage(pkg, err)
}

func lintPackage(pkgname string) int {
	pkg, err := build.Import(pkgname, "", 0)
	return lintImportedPackage(pkg, err)
}

func lintImportedPackage(pkg *build.Package, err error) int {
	if err != nil {
		if _, nogo := err.(*build.NoGoError); nogo {
			// Don't complain if the failure is due to no Go source files.
			return 0
		}
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	var files []string
	files = append(files, pkg.GoFiles...)
	files = append(files, pkg.TestGoFiles...)
	if pkg.Dir != "." {
		for i, f := range files {
			files[i] = filepath.Join(pkg.Dir, f)
		}
	}
	// TODO(dsymonds): Do foo_test too (pkg.XTestGoFiles)

	return lintFiles(files...)
}
