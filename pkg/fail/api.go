package fail

type ErrorFunc func(error) error

// Check reacts on non-nil err and runs error functions over it.
// It returns true if err != nil and false if err == nil.
func Check(err error, fns ...ErrorFunc) bool {
	if err == nil {
		return false
	}

	evaluate(err, fns)
	return true
}

// Recover react on non-nil recovered error and runs error funcsions over it.
// Does nothing if recovered is nil, but panics if recovered is not error.
func Recover(recovered any, fns ...ErrorFunc) {
	if err, ok := recovered.(error); ok {
		evaluate(err, fns)
		return
	}

	if recovered != nil {
		// if not an error panic again
		panic(recovered)
	}
}

// evaluate goes through all error functions
func evaluate(err error, efns []ErrorFunc) error {
	for _, fn := range efns {
		err = fn(err)
		if err == nil {
			break
		}
	}
	return err
}
