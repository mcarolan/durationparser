package durationparser

import "testing"

func TestParseDuration(t *testing.T) {
	cases := [] struct {
		Input string
		Result int64
	} {
		{ "1 second", 1000 },
		{ "2 seconds", 2000 },
		{ "1 hour 2 minutes and 1 second", 3721000 },	
		{ "1 hour 2 minutes 1 second", 3721000 },	
	}

	for _, c := range cases {
		r, err := ParseDuration(c.Input)

		if err != nil {
			t.Errorf("Unexpected error '%s' for input '%s'", err, c.Input)
		}
		if c.Result != r {
			t.Errorf("Expected %d actual %d", c.Result, r)
		}
	}
}
