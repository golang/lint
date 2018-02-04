// Test of stuttery names.

// Package donut ...
package donut

// Description ...
type DonutMaker struct{} // MATCH /donut\.DonutMaker.*stutter/

// Description ...
func DonutRank(d Donut) int { // MATCH /donut\.DonutRank.*stutter/
	return 0
}

// Description ...
type Donut struct{} // ok because it is the whole name

// Description ...
type Donuts []Donut // ok because it didn't start a new word

type donutGlaze int // ok because it is unexported

// Description ...
func (d *Donut) DonutMass() (grams int) { // okay because it is a method
	return 38
}
