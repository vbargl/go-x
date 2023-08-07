package flagutil

import (
	"flag"
	"strings"
)

func StringSliceFlag(addr *[]string) flag.Value {
	return (*stringSliceFlag)(addr)
}

type stringSliceFlag []string

func (s *stringSliceFlag) Set(v string) error {
	*s = append(*s, v)
	return nil
}

func (s *stringSliceFlag) String() string {
	return "[" + strings.Join(*s, ",") + "]"
}

var _ flag.Value = (*stringSliceFlag)(nil)
