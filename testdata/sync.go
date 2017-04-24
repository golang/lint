// Package pkgsync tests that exported types or methods from sync are flagged.
package pkgsync

import "sync"

type (
	// OK if unexported.
	b struct {
		sync.Mutex
		sync.Cond
		sync.Once
		sync.Pool
		sync.WaitGroup
	}

	// B is ...
	B struct { // MATCH /exported type B.*should.*expose.*sync\.Mutex.*or.*unexport/
		b int
		sync.Mutex
	}

	// C is ...
	C struct { // MATCH /exported type C.*should.*expose.*sync\.Cond.*or.*unexport/
		B B.b
		sync.Cond
	}

	// D is ...
	D struct { // MATCH /exported type D.*should.*expose.*sync\.Once.*or.*unexport/
		D int
		sync.Once
	}

	// E is ...
	E struct { // MATCH /exported type E.*should.*expose.*sync\.Pool.*or.*unexport/
		sync.Pool
	}

	// F is ...
	F struct { // MATCH /exported type F.*should.*expose.*sync\.WaitGroup.*or.*unexport/
		sync.WaitGroup
	}

	// G is ...
	G struct { // MATCH /exported type G.*should.*expose.*sync\.RWMutex.*or.*unexport/
		MU sync.RWMutex
	}
)
