// Package gopragmas test the go:... pragmas
package gopragmas

//go:generate ok         OK
// go:generate not ok    MATCH /go: pragmas should have no spaces/
