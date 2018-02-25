// MATCH:0 /at least one file should have a valid package comment/
// Test of block package comment with leading space.

/*
 Package foo is pretty sweet.
MATCH /package comment.*leading space/
*/
package foo
