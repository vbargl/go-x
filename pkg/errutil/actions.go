package errutil

import (
	"fmt"
	"log"
)

func HandleIfRecovered(recovered any, handler func(err error)) {
	if err, ok := recovered.(error); ok {
		handler(err)
	}
}

func LogFatalIfErr(err error, wrapper ErrorWrapper) bool {
	if err == nil {
		return false
	}

	log.Fatal(wrapper.wrap(err))
	return true
}

func LogPrintIfErr(err error, wrapper ErrorWrapper) bool {
	if err == nil {
		return false
	}

	log.Print(wrapper.wrap(err))
	return true
}

func PrintIfErr(err error, wrapper ErrorWrapper) bool {
	if err == nil {
		return false
	}

	fmt.Println(wrapper.wrap(err))
	return true
}

func AssignIfErr(target *error, err error, wrapper ErrorWrapper) bool {
	if err == nil {
		return false
	}

	*target = wrapper.wrap(err)
	return true
}
