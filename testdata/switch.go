// Test for switch case statements.

// Package pkg ...
package pkg

import "fmt"

func bad() {
	a := 0

	for {
		switch 1 {
		case 0: // MATCH /are you missing a fallthrough?/
		case 1, 2: // MATCH /are you missing a fallthrough?/
		case 9: // MATCH /statement does no work/
			{
			}
		case 10: // MATCH /statement does no work/
			{
				{
				}
			}
		case 11: // MATCH /statement does no work/
			if a == 0 {
			}
		case 12: // MATCH /statement does no work/
			for {
			}
		default: // MATCH /statement does no work/

		}

		break
	}

	switch 1 {
	default: // MATCH /are you missing a fallthrough?/
	case 1: // MATCH /statement does no work/
	}

}

func nested() {
	var b int
	switch 1 {

	case 1:
		switch 'a' {
		case 'a': // MATCH /are you missing a fallthrough?/
		default:
			b = 9
		}

	case 2: // MATCH /statement does no work/
		switch 'a' {
		case 'a': // MATCH /are you missing a fallthrough?/
		default: // MATCH /statement does no work/
		}

	case 3:
		f()
	}
}

func good() {
	var a int
	for {
		switch 0 {
		case 0:
			fallthrough
		case 1:
			return
		case 2:
			break
		case 3:
			continue
		case 4:
			fmt.Println(a)
		case 5:
			a = 5
		case 6:
			f() // one could argue this isn't actually doing anything
		case 7:
			if z := len("hello"); z > 0 {
				fmt.Println("hello")
			}
		case 8:
			go f()
		case f():
		default:
			break
		}
	}
}

func f() {
}
