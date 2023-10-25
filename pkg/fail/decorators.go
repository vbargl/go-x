package fail

import "fmt"

// Combine merges all fns into single [barglvojtech.net/x/fail.ErrorRefiner]
func Combine(ers ...ErrorRefiner) ErrorRefiner {
	c := make(errorCombiner, 0, len(ers))
	for _, er := range ers {
		switch er := er.(type) {
		case nil:
			continue
		case errorCombiner:
			c = append(c, er...)
		default:
			c = append(c, er)
		}
	}

	switch len(c) {
	case 0:
		panic("fail: no usable ErrorRefiner")
	case 1:
		return c[0]
	default:
		return c
	}
}

type errorCombiner []ErrorRefiner

func (ers errorCombiner) RefineError(err error) error {
	return evaluate(err, ers)
}

// Wrap formats error in given format using fmt.Errorf function.
// Format must contain '%w' to wrap given error.
func Wrap(format string) ErrorRefinerFunc {
	return func(err error) error {
		return fmt.Errorf(format, err)
	}
}

// Wrapf formats error in given format using fmt.Errorf function.
// Format must contain '%%w' to wrap given error.
func Wrapf(format string, args ...any) ErrorRefinerFunc {
	return func(err error) error {
		format := fmt.Sprintf(format, args...)
		return fmt.Errorf(format, err)
	}
}
