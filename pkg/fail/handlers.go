package fail

import (
	"fmt"
	"log"
)

// Assign assigns error into target.
// Returns err.
func Assign(target *error) ErrorFunc {
	return func(err error) error {
		*target = err
		return err
	}
}

// Panic panics when evaluated.
// Does not return.
func Panic() ErrorFunc {
	return func(err error) error {
		panic(err)
	}
}

// Break just retunrs nil and stoping it from furhter evaluation of error functions.
func Break() ErrorFunc {
	return func(err error) error {
		return nil
	}
}

// LogPrint prints error using Log.Print.
func LogPrint() ErrorFunc {
	return func(err error) error {
		log.Print(err)
		return err
	}
}

// LogPanic prints error using Log.Panic which panics right after evaluation.
func LogPanic() ErrorFunc {
	return func(err error) error {
		log.Panic(err)
		return nil
	}
}

// LogPrint prints error using Log.Fatal which exists program right after evaluation.
func LogFatal() ErrorFunc {
	return func(err error) error {
		log.Fatal(err)
		return nil
	}
}

// LoggerPrint prints error using Logger.Print.
func LoggerPrint(logger *log.Logger) ErrorFunc {
	return func(err error) error {
		logger.Print(err)
		return nil
	}
}

// LoggerPanic prints error using Logger.Panic which panics right after evaluation.
func LoggerFatal(logger *log.Logger) ErrorFunc {
	return func(err error) error {
		logger.Fatal(err)
		return nil
	}
}

// LoggerPrint prints error using Logger.Fatal which exists program right after evaluation.
func LoggerPanic(logger *log.Logger) ErrorFunc {
	return func(err error) error {
		logger.Panic(err)
		return nil
	}
}

// Print prints error using fmt.Prinrln.
// returns err
func Print() ErrorFunc {
	return func(err error) error {
		fmt.Println(err)
		return err
	}
}
