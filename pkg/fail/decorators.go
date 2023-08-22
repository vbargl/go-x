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

// Prefix prefixes error in format "<prefix>: <original error message>" using fmt.Errorf function.
func Prefix(prefix string) ErrorRefinerFunc {
	return func(err error) error {
		return fmt.Errorf("%s: %v", prefix, err)
	}
}

// Prefixf prefixes error in format "<prefix>: <original error message>" using fmt.Errorf function.
// Prefix is formated string using fmt.Sprintf function by passing format and args.
func Prefixf(format string, args ...any) ErrorRefinerFunc {
	return func(err error) error {
		prefix := fmt.Sprintf(format, args...)
		return fmt.Errorf("%s: %v", prefix, err)
	}
}
