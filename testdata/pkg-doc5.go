// MATCH:0 /at least one file should have a valid package comment/
// Test of detached package comment.

/*
Package foo is pretty sweet.
*/

package foo

// MATCH:7 /package comment.*detached/
