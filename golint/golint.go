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
	"strings"

	"github.com/golang/lint"
)

var (
	minConfidence = flag.Float64("min_confidence", 0.8, "minimum confidence of a problem to print it")
	setExitStatus = flag.Bool("set_exit_status", false, "set exit status to 1 if any issues are found")
	suggestions   int
)

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

	var foundErrors bool

	switch flag.NArg() {
	case 0:
		foundErrors = lintDir(".")
	case 1:
		arg := flag.Arg(0)
		if strings.HasSuffix(arg, "/...") && isDir(arg[:len(arg)-4]) {
			for _, dirname := range allPackagesInFS(arg) {
				foundErrors = foundErrors || lintDir(dirname)
			}
		} else if isDir(arg) {
			foundErrors = lintDir(arg)
		} else if exists(arg) {
			foundErrors = lintFiles(arg)
		} else {
			for _, pkgname := range importPaths([]string{arg}) {
				foundErrors = foundErrors || lintPackage(pkgname)
			}
		}
	default:
		foundErrors = lintFiles(flag.Args()...)
	}

	if foundErrors {
		os.Exit(1)
	}

	if *setExitStatus && suggestions > 0 {
		fmt.Fprintf(os.Stderr, "Found %d lint suggestions; failing.\n", suggestions)
		os.Exit(1)
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

func lintFiles(filenames ...string) (foundErrors bool) {
	files := make(map[string][]byte)
	for _, filename := range filenames {
		src, err := ioutil.ReadFile(filename)
		if err != nil {
			foundErrors = true
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		files[filename] = src
	}

	l := new(lint.Linter)
	ps, err := l.LintFiles(files)
	if err != nil {
		foundErrors = true
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	for _, p := range ps {
		if p.Confidence >= *minConfidence {
			foundErrors = true
			fmt.Printf("%v: %s\n", p.Position, p.Text)
			suggestions++
		}
	}

	return foundErrors
}

func lintDir(dirname string) bool {
	pkg, err := build.ImportDir(dirname, 0)
	return lintImportedPackage(pkg, err)
}

func lintPackage(pkgname string) bool {
	pkg, err := build.Import(pkgname, ".", 0)
	return lintImportedPackage(pkg, err)
}

func lintImportedPackage(pkg *build.Package, err error) (foundErrors bool) {
	if err != nil {
		if _, nogo := err.(*build.NoGoError); nogo {
			// Don't complain if the failure is due to no Go source files.
			return
		}
		foundErrors = true
		fmt.Fprintln(os.Stderr, err)
		return
	}

	var files []string
	files = append(files, pkg.GoFiles...)
	files = append(files, pkg.CgoFiles...)
	files = append(files, pkg.TestGoFiles...)
	if pkg.Dir != "." {
		for i, f := range files {
			files[i] = filepath.Join(pkg.Dir, f)
		}
	}
	// TODO(dsymonds): Do foo_test too (pkg.XTestGoFiles)

	return foundErrors || lintFiles(files...)
}
