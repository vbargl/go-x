package fail

import "fmt"

// Combine merges all fns into single ErrorFunc
func Combine(fns ...ErrorFunc) ErrorFunc {
	return func(err error) error {
		return evaluate(err, fns)
	}
}

// Prefix prefixes error in format "<prefix>: <original error message>" using fmt.Errorf function.
func Prefix(prefix string) ErrorFunc {
	return func(err error) error {
		return fmt.Errorf("%s: %v", prefix, err)
	}
}

// Prefixf prefixes error in format "<prefix>: <original error message>" using fmt.Errorf function.
// Prefix is formated string using fmt.Sprintf function by passing format and args.
func Prefixf(format string, args ...any) ErrorFunc {
	return func(err error) error {
		prefix := fmt.Sprintf(format, args...)
		return fmt.Errorf("%s: %v", prefix, err)
	}
}
