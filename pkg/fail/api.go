package fail

type ErrorRefiner interface {
	RefineError(error) error
}

type ErrorRefinerFunc func(error) error

func (fn ErrorRefinerFunc) RefineError(err error) error {
	return fn(err)
}

// Check reacts on non-nil err and runs error functions over it.
// It returns true if err != nil and false if err == nil.
func Check(err error, ers ...ErrorRefiner) bool {
	if err == nil {
		return false
	}

	evaluate(err, ers)
	return true
}

// Recover react on non-nil recovered error and runs error funcsions over it.
// Does nothing if recovered is nil, but panics if recovered is not error.
func Recover(ers ...ErrorRefiner) {
	recovered := recover()
	if err, ok := recovered.(error); ok {
		evaluate(err, ers)
		return
	}

	if recovered != nil {
		// if not an error panic again
		panic(recovered)
	}
}

// evaluate goes through all error functions
func evaluate(err error, ers []ErrorRefiner) error {
	for _, er := range ers {
		if er == nil {
			continue
		}

		err = er.RefineError(err)
		if err == nil {
			break
		}
	}
	return err
}
