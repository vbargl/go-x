package stringutil

import (
	"fmt"
	"strings"
)

// Interpolate will replace all variables formatted ${varname} in s
// with values provided by variables map.
//
// Variables must be formatted as ident token.
// See test file for examples.
func Interpolate(s string, variables map[string]string) string {
	for variable, value := range variables {
		s = strings.ReplaceAll(s, fmt.Sprintf("${%s}", variable), value)
	}
	return s
}
