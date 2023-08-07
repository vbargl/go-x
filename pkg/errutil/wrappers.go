package errutil

import "fmt"

func PrefixWith(prefix string) ErrorWrapper {
	return func(err error) error {
		return fmt.Errorf("%s: %v", prefix, err)
	}
}

func PrefixWithFormatted(format string, args ...any) ErrorWrapper {
	return func(err error) error {
		prefix := fmt.Sprintf(format, args...)
		return fmt.Errorf("%s: %v", prefix, err)
	}
}
