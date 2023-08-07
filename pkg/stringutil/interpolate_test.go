package stringutil

import (
	"fmt"
	"testing"
)

func TestInterpolate(t *testing.T) {
	type param struct {
		format    string
		variables map[string]string
		expected  string
	}

	params := []param{
		{
			format:    "Hello ${name}!",
			variables: map[string]string{"name": "World"},
			expected:  "Hello World!",
		},
		{
			format: "http://example.com/path/${var1}/path/${var2}/path",
			variables: map[string]string{
				"var1": "value1",
				"var2": "value2",
			},
			expected: "http://example.com/path/value1/path/value2/path",
		},
		{
			format: "${verb} ${verb} ${verb}, ${name}!",
			variables: map[string]string{
				"verb": "knock",
				"name": "Penny",
			},
			expected: "knock knock knock, Penny!",
		},
	}

	for _, p := range params {
		t.Run(fmt.Sprintf("formating %s", p.format), func(t *testing.T) {
			if got := Interpolate(p.format, p.variables); got != p.expected {
				t.Errorf("expected %s, got %s", p.expected, got)
			}
		})
	}
}
