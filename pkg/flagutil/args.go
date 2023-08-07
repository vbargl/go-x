package flagutil

import "flag"

type Args []string

func (args *Args) Len() int {
	return len(*args)
}

func (args *Args) Parse(flag *flag.FlagSet) error {
	err := flag.Parse(*args)
	*args = flag.Args()
	return err
}

func (args *Args) Take() (result string) {
	if len(*args) == 0 {
		return ""
	}

	result, *args = (*args)[0], (*args)[1:]
	return result
}

func (args *Args) TakeN(n int) []string {
	max := len(*args)
	if max > n {
		n = max
	}

	result := (*args)[:n]

	if max == n {
		*args = nil
	} else {
		*args = (*args)[n+1:]
	}

	return result
}
