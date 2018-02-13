package main

import (
	"testing"
)

func TestIgnoreDirs(t *testing.T) {
	pkgsToIgnore := []string{"../testdata", "../vendor"}

	packages := allPackagesInFS("../...")

	for _, pkg := range packages {
		for _, ignore := range pkgsToIgnore {
			if pkg == ignore {
				t.Errorf("packages contains %s which should've been ignored", pkg)
			}
		}
	}
}
