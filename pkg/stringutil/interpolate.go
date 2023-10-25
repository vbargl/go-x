package stringutil

import (
	"fmt"
	"strings"
)

// InterpolateMap will replace all variables formatted ${varname} in s
// with values provided by variables map.
// Interpolated value can be [string], [fmt.Stringer] or [func() string]
//
// Variables must be formatted as ident token.
// See test file for examples.
func InterpolateMap(s string, variables map[string]any) string {
	return interpolate(s, func(v string) string {
		value, ok := variables[v]
		if !ok {
			return fmt.Sprintf("${%s}", v)
		}

		switch value := value.(type) {
		case string:
			return value
		case fmt.Stringer:
			return value.String()
		case func() string:
			return value()
		default:
			return fmt.Sprintf("${%s}", v)
		}
	})
}

// InterpolateFunc replace each found variable formatted ${varname} through given fn function.
//
// Variables must be formatted as ident token.
// See test file for examples.
func InterpolateFunc(s string, fn func(string) string) string {
	return interpolate(s, fn)
}

func interpolate(s string, fn func(string) string) string {
	var (
		sb             strings.Builder
		vars           = make([]finding, 0, 24)
		offset, length = 0, len(s)
	)

	for {
		start := strings.IndexRune(s[offset:], '$')
		if start == -1 {
			break
		}

		// start was just offset from startIndex till this point,
		// add startIndex to represent offset from 0.
		start += offset
		if start+1 >= length {
			break
		}

		switch r := s[start+1]; r {

		default: // '$<any other then $ or {>' - ignore
			offset += 1

		case '$': // '$$' - escaped $
			vars = append(vars, finding{
				findingType: f_Dollar,
				startIndex:  start,
				endIndex:    start + 2,
			})

			offset = start + 2

		case '{': // '${' - start
			end := strings.IndexRune(s[start+2:], '}')
			if end == -1 {
				offset = length
				break // '${var' at the end - ignore
			}

			vars = append(vars, finding{
				findingType: f_Var,
				startIndex:  start,                    // index of '$' rune
				endIndex:    start + 2 + end + 1,      // index after '}' rune
				name:        s[start+2 : start+2+end], // seq between '${' and '}'
			})

			offset = start + 2 + end + 1
		}
	}

	offset = 0
	for _, v := range vars {
		sb.WriteString(s[offset:v.startIndex])
		switch v.findingType {
		case f_Dollar:
			sb.WriteRune('$')
		case f_Var:
			value := fn(v.name)
			sb.WriteString(value)
		}

		offset = v.endIndex
	}
	sb.WriteString(s[offset:])

	return sb.String()
}

type findingType int

const (
	f_Dollar findingType = iota
	f_Var
)

type finding struct {
	findingType          findingType
	startIndex, endIndex int

	// f_Var
	name string
}
