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
	"os"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/tools/go/loader"

	"github.com/golang/lint"
)

var minConfidence = flag.Float64("min_confidence", 0.8, "minimum confidence of a problem to print it")
var buildTags stringsFlag

func init() {
	flag.Var(&buildTags, "tags", "a list of build tags to consider satisfied during the build")
}

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

	ctx := &build.Default
	for _, tag := range buildTags {
		ctx.BuildTags = append(ctx.BuildTags, tag)
	}

	conf := &loader.Config{
		Build:      ctx,
		ParserMode: parser.ParseComments,
	}

	switch flag.NArg() {
	case 0:
		addDir(conf, ".")
	case 1:
		arg := flag.Arg(0)
		if strings.HasSuffix(arg, "/...") && isDir(arg[:len(arg)-4]) {
			for _, dirname := range allPackagesInFS(arg) {
				addDir(conf, dirname)
			}
		} else if isDir(arg) {
			addDir(conf, arg)
		} else if exists(arg) {
			conf.CreateFromFilenames(".", arg)
		} else {
			for _, pkgname := range importPaths([]string{arg}, buildTags) {
				conf.ImportWithTests(pkgname)
			}
		}
	default:
		conf.CreateFromFilenames(".", flag.Args()...)
	}

	prog, err := conf.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	l := new(lint.Linter)
	var ps []lint.Problem

	for _, pkg := range prog.InitialPackages() {
		pp, err := l.LintFiles(prog.Fset, pkg.Files)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}

		ps = append(ps, pp...)
	}

	sort.Sort(lint.ByPosition(ps))

	for _, p := range ps {
		if p.Confidence >= *minConfidence {
			fmt.Printf("%v: %s\n", p.Position, p.Text)
		}
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

func addDir(conf *loader.Config, dirname string) {
	// go/loader does currently not expose ImportDir
	pkg, err := build.ImportDir(dirname, 0)
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

	joinDirWithFilenames(dirname, files)

	conf.CreateFromFilenames(".", files...)

	if files := pkg.XTestGoFiles; len(files) != 0 {
		joinDirWithFilenames(dirname, files)

		conf.CreateFromFilenames(".", files...)
	}
}

func joinDirWithFilenames(dir string, files []string) {
	if dir != "." {
		for i, f := range files {
			files[i] = filepath.Join(dir, f)
		}
	}
}
