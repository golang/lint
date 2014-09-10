// Test for variable capturing in closure.

// Package foo ...
package foo

func test() {
	for i := 0; i < 128; i++ {
		go func() {
			_ = i // MATCH /closure should not capture loop variable i/
		}()
	}

	for i := 0; i < 128; i += 2 {
		go func() {
			_ = i // MATCH /closure should not capture loop variable i/
		}()
	}

	for key, value := range map[int]int{} {
		go func() {
			_ = key   // MATCH /closure should not capture loop variable key/
			_ = value // MATCH /closure should not capture loop variable value/
		}()
	}

	for _, value := range map[int]int{} {
		go func() {
			_ = value // MATCH /closure should not capture loop variable value/
		}()
	}

	for key := range map[int]int{} {
		go func() {
			_ = key // MATCH /closure should not capture loop variable key/
		}()
	}
}
