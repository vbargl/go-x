package sliceutil

// Iterate iterates through all given fns functions and pass given v value.
// Similar to [sliceutil.IterateAll].
func Iterate[T any, Fn ~func(T)](v T, fns ...Fn) {
	for _, fn := range fns {
		fn(v)
	}
}

// IterateAll iterates through all given fns functions and pass given v value.
// Similar to [sliceutil.Iterate].
func IterateAll[T any, Fn ~func(T)](v T, fns []Fn) {
	for _, fn := range fns {
		fn(v)
	}
}
