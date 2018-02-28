// Test that adjacent function arguments of the same type only have a single type declaration.

// Package pkg ...
package pkg

import "net/http"

var x func(x int, y int) // MATCH /adjacent arguments of like type should have a single type identifier/
var x2 func(x, y int)    // should be ok

func a(s []*int, r []*int, p string) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func az(s []*int, r []*int, p string) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func a2(s []int, r []int, p string) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func b(s []bool, r []bool, p string) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func b2(s []*bool, r []*bool, p string) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func c(s string, r string, p string) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func c2(s *string, r *string, p string) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func d(s int, r int, p string) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func d2(s *int, r *int, p string) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func e(s int16, r int16, p string) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func e2(s *int16, r *int16, p string) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func f(s int32, r int32, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func f2(s *int32, r *int32, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func g(s int64, r int64, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func g2(s *int64, r *int64, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func h(s interface{}, r interface{}, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func h2(s *interface{}, r *interface{}, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func h3(s interface{}, r interface { // should be ok
	H()
}) {
}
func i(s struct{}, r struct{}, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func i2(a *struct{}, r *struct{}, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func i3(s struct{ a, b int }, t struct { // MATCH /adjacent arguments of like type should have a single type identifier/
	a,
	b int
}) {
}
func i4(s struct {
	T
	a int
}, t struct {
	t T
	a int
}) {
}
func j(s chan *int, r chan *int, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func j2(s chan *string, r chan *string, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func j3(s chan int, r chan int, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func j4(s *chan int, r *chan int, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func k(s [9]int, r [9]int, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func k2(s []string, r []string, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func k3(s []*string, r []*string, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func l(s func(a int, b int), r func(a int, b int), p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func l2(s *func(a string, b int), r *func(a string, b int), p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func l3(s func(a int), r func(a, b int), p int16) { // should be ok
}
func l4(s func(a, b int), t func(b, c int)) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func m(s rune, r rune, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func m2(s *rune, r *rune, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func n(s float32, r float32, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func n2(s *float32, r *float32, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func o(s float64, r float64, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func o2(s *float64, r *float64, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func p(s complex64, r complex64, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func p2(s *complex64, r *complex64, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func q(s complex128, r complex128, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func q2(s *complex128, r *complex128, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func r(a uintptr, r uintptr, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func r2(a *uintptr, r *uintptr, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func s(a error, r error, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func s2(a error, r *error, p int16) { // should be ok
}
func t(a uint, r uint, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func t2(a *uint, r *uint, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func u(a uint8, r uint8, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func u2(a *uint8, r *uint8, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func v(a uint16, r uint16, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func v2(a *uint16, r *uint16, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func w(a uint32, r uint32, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func w2(a *uint32, r *uint32, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func x(a *uint64, r *uint64, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func y(a bool, r bool, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func y2(a *bool, r *bool, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func z(a map[string]int, r map[string]int, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func z2(a map[*string]int, r map[string]*int, p int16) { // should be ok
}
func z3(a map[*string]*int, r map[*string]int, p int16) { // should be ok
}
func z4(a map[*string]int, r map[*string]int, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func z5(a *map[*string]int, r *map[*string]*int, p int16) { // should be ok
}

func z6(a, b *http.Request) { // should be ok
}
func z6(a *http.Request, b *http.Request) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func z8(a func(...string), b func(...string)) { // MATCH /adjacent arguments of like type should have a single type identifier/
}

// T ...
type T struct{}

func aa(a T, r T, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
func ab(a *T, r *T, p int16) { // MATCH /adjacent arguments of like type should have a single type identifier/
}

// TI ...
type TI interface{}

func ba(a TI, b TI) { // MATCH /adjacent arguments of like type should have a single type identifier/
}
