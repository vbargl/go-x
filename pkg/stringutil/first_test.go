package stringutil

import "testing"

func TestFirstNonEmpty(t *testing.T) {
	type param struct {
		given    []string
		expected string
	}

	params := []param{
		{given: []string{"", "", ""}, expected: ""},
		{given: []string{"", "something"}, expected: "something"},
		{given: []string{"", "", "", "something"}, expected: "something"},
	}

	for i, p := range params {
		if got := FirstNonEmpty(p.given...); got != p.expected {
			t.Errorf("expected %s, got %s (case %d)", p.expected, got, i)
		}
	}
}
