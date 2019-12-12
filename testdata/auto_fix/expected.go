// Test for name linting.

// Package pkg_with_underscores ...
package pkgWithUnderscores // MATCH /underscore.*package name/

import (
	"io"
	"net"
	netHTTP "net/http" // renamed deliberately
	"net/url"
)

import "C"

var varName int // MATCH /underscore.*var.*var_name/

type tWow struct { // MATCH /underscore.*type.*t_wow/
	xDamn int      // MATCH /underscore.*field.*x_damn/
	URL   *url.URL // MATCH /struct field.*Url.*URL/
}

const fooID = "blah" // MATCH /fooId.*fooID/

func fIt() { // MATCH /underscore.*func.*f_it/
	moreUnderscore := 4 // MATCH /underscore.*var.*more_underscore/
	_ = moreUnderscore
	var err error
	if isEOF := (err == io.EOF); isEOF { // MATCH /var.*isEof.*isEOF/
		moreUnderscore = 7 // should be okay
	}

	x := netHTTP.Request{} // should be okay
	_ = x

	var ips []net.IP
	for _, theIP := range ips { // MATCH /range var.*theIp.*theIP/
		_ = theIP
	}

	switch myJSON := g(); { // MATCH /var.*myJson.*myJSON/
	default:
		_ = myJSON
	}
	var y netHTTP.ResponseWriter // an interface
	switch tAPI := y.(type) {    // MATCH /var.*tApi.*tAPI/
	default:
		_ = tAPI
	}

	var c chan int
	select {
	case qID := <-c: // MATCH /var.*qId.*qID/
		_ = qID
	}
}

// Common styles in other languages that don't belong in Go.
const (
	CPP_CONST  = 1 // MATCH /ALL_CAPS.*CamelCase/
	leadingKay = 2 // MATCH /k.*leadingKay/

	HTML    = 3 // okay; no underscore
	X509B   = 4 // ditto
	V1_10_5 = 5 // okay; fewer than two uppercase letters
)

var varsAreSometimesUsedAsConstants = 0 // MATCH /k.*varsAreSometimesUsedAsConstants/
var (
	varsAreSometimesUsedAsConstants2 = 0 // MATCH /k.*varsAreSometimesUsedAsConstants2/
)

var thisIsNotOkay = struct { // MATCH /k.*thisIsNotOkay/
	thisIsOkay bool
}{}

func thisIsOkay() { // this is okay because this is a function name
	var thisIsAlsoOkay = 1 // this is okay because this is a non-top-level variable
	_ = thisIsAlsoOkay
	const thisIsNotOkay = 2 // MATCH /k.*thisIsNotOkay/
}

var anotherFunctionScope = func() {
	var thisIsOkay = 1 // this is okay because this is a non-top-level variable
	_ = thisIsOkay
	const thisIsNotOkay = 2 // MATCH /k.*thisIsNotOkay/}
}

func f(badName int)                   {}            // MATCH /underscore.*func parameter.*bad_name/
func g() (noWay int)                  { return 0 }  // MATCH /underscore.*func result.*no_way/
func (t *tWow) f(moreUnder string)    {}            // MATCH /underscore.*method parameter.*more_under/
func (t *tWow) g() (stillMore string) { return "" } // MATCH /underscore.*method result.*still_more/

type i interface {
	CheckHTML() string // okay; interface method names are often constrained by the concrete types' method names

	F(fooBar int) // MATCH /foo_bar.*fooBar/
}

// All okay; underscore between digits
const case1_1 = 1

type case2_1 struct {
	case2_2 int
}

func case3_1(case3_2 int) (case3_3 string) {
	case3_4 := 4
	_ = case3_4

	return ""
}

type t struct{}

func (t) LastInsertId() (int64, error) { return 0, nil } // okay because it matches a known style violation

//export exported_to_c
func exportedToC() {} // okay: https://github.com/golang/lint/issues/144

//export exported_to_c_with_arg
func exportedToCWithArg(butUseGoParamNames int) // MATCH /underscore.*func parameter.*but_use_go_param_names/

// This is an exported C function with a leading doc comment.
//
//export exported_to_c_with_comment
func exportedToCWithComment() {} // okay: https://github.com/golang/lint/issues/144

//export maybe_exported_to_CPlusPlusWithCamelCase
func maybeExportedToCPlusPlusWithCamelCase() {} // okay: https://github.com/golang/lint/issues/144

// WhyAreYouUsingCapitalLetters_InACFunctionName is a Go-exported function that
// is also exported to C as a name with underscores.
//
// Don't do that. If you want to use a C-style name for a C export, make it
// lower-case and leave it out of the Go-exported API.
//
//export WhyAreYouUsingCapitalLetters_InACFunctionName
func WhyAreYouUsingCapitalLettersInACFunctionName() {} // MATCH /underscore.*func.*Why.*CFunctionName/
