package ioutil

import "time"

type Interrupter interface {
	// SetDeadline sets the read and write deadlines for a File.
	// It is equivalent to calling both SetReadDeadline and SetWriteDeadline.
	//
	// Only some kinds of files support setting a deadline. Calls to SetDeadline
	// for files that do not support deadlines will return ErrNoDeadline.
	// On most systems ordinary files do not support deadlines, but pipes do.
	//
	// A deadline is an absolute time after which I/O operations fail with an
	// error instead of blocking. The deadline applies to all future and pending
	// I/O, not just the immediately following call to Read or Write.
	// After a deadline has been exceeded, the connection can be refreshed
	// by setting a deadline in the future.
	//
	// If the deadline is exceeded a call to Read or Write or to other I/O
	// methods will return an error that wraps ErrDeadlineExceeded.
	// This can be tested using errors.Is(err, os.ErrDeadlineExceeded).
	// That error implements the Timeout method, and calling the Timeout
	// method will return true, but there are other possible errors for which
	// the Timeout will return true even if the deadline has not been exceeded.
	//
	// An idle timeout can be implemented by repeatedly extending
	// the deadline after successful Read or Write calls.
	//
	// A zero value for t means I/O operations will not time out.
	SetDeadline(time.Time) error
}

type InterruptReader interface {
	Interrupter

	// SetReadDeadline sets the deadline for future Read calls and any
	// currently-blocked Read call.
	// A zero value for t means Read will not time out.
	// Not all files support setting deadlines; see SetDeadline.
	SetReadDeadline(time.Time) error
}

type InterruptWriter interface {
	Interrupter

	// SetWriteDeadline sets the deadline for any future Write calls and any
	// currently-blocked Write call.
	// Even if Write times out, it may return n > 0, indicating that
	// some of the data was successfully written.
	// A zero value for t means Write will not time out.
	// Not all files support setting deadlines; see SetDeadline.
	SetWriteDeadline(time.Time) error
}
