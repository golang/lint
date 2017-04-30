// Package gogeneratedirective test the go:generate directive
package gogeneratedirective

//go:generate ok         OK
// go:generate not ok    MATCH /go:generate directive should have no spaces/
