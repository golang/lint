// Package foo ...
package foo

var ch1, ch2 chan int

func fn1() {
	select { // MATCH /simple channel send/
	case <-ch1:
	}

	select {
	case <-ch1:
	case <-ch2:
	}

	select {
	case <-ch1:
	default:
	}
}

func fn2() {
	for { // MATCH /range instead of for/
		select {
		case <-ch1:
		}
	}
}

func fn3() {
	for {
		select {
		case <-ch1:
		case <-ch2:
		}
	}
}

func fn4() {
	for {
		select { // MATCH /simple channel send/
		case <-ch1:
		}
		break
	}
}
