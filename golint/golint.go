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
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

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
		lintDir(".")
	case 1:
		arg := flag.Arg(0)
		if strings.HasSuffix(arg, "/...") && isDir(arg[:len(arg)-4]) {
			for _, dirname := range allPackagesInFS(arg) {
				lintDir(dirname)
			}
		} else if isDir(arg) {
			lintDir(arg)
		} else if exists(arg) {
			lintFiles(arg)
		} else {
			for _, pkgname := range importPaths([]string{arg}) {
				lintPackage(pkgname)
			}
		}
	default:
		lintFiles(flag.Args()...)
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

func lintFiles(filenames ...string) {
	files := make(map[string]struct{})         // Map of already loaded files
	pkgs := make(map[string]map[string][]byte) // Map of packages each holding a map of their loaded files
	for _, filename := range filenames {
		if _, ok := files[filename]; ok {
			continue
		}

		src, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		files[filename] = struct{}{}

		pkgName := path.Dir(filename)
		if strings.HasSuffix(filename, "_test.go") {
			f, err := parser.ParseFile(token.NewFileSet(), "", src, 0)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				return
			}

			if n := f.Name.Name; strings.HasSuffix(n, "_test") {
				pkgName = pkgName + " " + n
			}
		}

		pkg, ok := pkgs[pkgName]
		if !ok {
			pkg = make(map[string][]byte)

			pkgs[pkgName] = pkg
		}

		pkg[filename] = src
	}

	pkgsNames := make([]string, 0, len(pkgs))
	for n := range pkgs {
		pkgsNames = append(pkgsNames, n)
	}
	sort.Strings(pkgsNames)

	l := new(lint.Linter)

	for _, n := range pkgsNames {
		ps, err := l.LintFiles(pkgs[n])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return
		}
		for _, p := range ps {
			if p.Confidence >= *minConfidence {
				fmt.Printf("%v: %s\n", p.Position, p.Text)
			}
		}
	}
}

func lintDir(dirname string) {
	pkg, err := build.ImportDir(dirname, 0)
	lintImportedPackage(pkg, err)
}

func lintPackage(pkgname string) {
	pkg, err := build.Import(pkgname, ".", 0)
	lintImportedPackage(pkg, err)
}

func lintImportedPackage(pkg *build.Package, err error) {
	if err != nil {
		if _, nogo := err.(*build.NoGoError); nogo {
			// Don't complain if the failure is due to no Go source files.
			return
		}
		fmt.Fprintln(os.Stderr, err)
		return
	}

	var files []string
	files = append(files, pkg.GoFiles...)
	files = append(files, pkg.TestGoFiles...)
	joinDirWithFilenames(pkg.Dir, files)

	lintFiles(files...)

	if len(pkg.XTestGoFiles) > 0 {
		files := pkg.XTestGoFiles
		joinDirWithFilenames(pkg.Dir, files)

		lintFiles(files...)
	}
}

func joinDirWithFilenames(dir string, files []string) {
	if dir != "." {
		for i, f := range files {
			files[i] = filepath.Join(dir, f)
		}
	}
}
