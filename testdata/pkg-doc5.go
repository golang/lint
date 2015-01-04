// Test of detached package comment.

/*
Package foo is pretty sweet.
*/

package foo // MATCH /package comment.*detached/
