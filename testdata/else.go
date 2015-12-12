// Test of return+else warning.

// Package pkg ...
package pkg

import "log"

func e(x int) bool {
	if x > 0 {
		return true
	}
	return false
}

func f(x int) bool {
	if x > 0 {
		return true
	} else { // MATCH /if.*return.*else.*outdent/
		log.Printf("non-positive x: %d", x)
	}
	return false
}

func g(f func() bool) string {
	if ok := f(); ok {
		return "it's okay"
	} else { // MATCH /if.*return.*else.*outdent.*short.*var.*declaration/
		return "it's NOT okay!"
	}
}

func h() bool {
	if false {
		return false
	} else if false {
		return false
	} else {
		return true
	}
}

func i() bool {
	if false {
		return false
	} else if false {
		return false
	} else if false {
		return false
	} else {
		return true
	}
}

func j() bool {
	if false {
	} else if false {
	} else if false {
		return false
	} else {
	}
	return true
}

func k() bool {
	if false {
		return false
	} else if false {
		return false
	} else if false {
		return false
	} else {
	}
	return true
}
