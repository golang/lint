// Test that we ask for the right kinds of comments for types.

// Package pkg ...
package pkg

type TypeA []int // MATCH /should have comment/

// TypeB ...
type TypeB []int

// A TypeC ...
type TypeC []int

// An ItemBeginningWithAVowel ...
type ItemBeginningWithAVowel []int

// Type TypeD ...
type TypeD []int

/* MATCH /comment on exported type TypeE/ */ // Much ado about TypeE ...
type TypeE []int
