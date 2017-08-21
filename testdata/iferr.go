// Test of redundant if err != nil

// Package pkg ...
package pkg

func f() error {
	if err := f(); err != nil {
		g()
		return err
	}
	return nil
}

func g() error {
	if err := f(); err != nil { // MATCH /redundant/
		return err
	}
	return nil
}

func h() error {
	if err, x := f(), 1; err != nil {
		return err
	}
	return nil
}

func i() error {
	a := 1
	if err := f(); err != nil {
		a++
		return err
	}
	return nil
}

func multi() error {
	a := 0
	var err error
	if true {
		a++
		if err := f(); err != nil { // MATCH /redundant/
			return err
		}
		return nil
		a++
	} else {
		a++
		if err = f(); err != nil { // MATCH /redundant/
			return err
		}
		return nil
		a++
	}
}
