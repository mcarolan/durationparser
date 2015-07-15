package durationparser

import (
	"testing"
	"errors"
)

func TestParseDuration(t *testing.T) {
	cases := [] struct {
		Input string
		Result int64
		Error error
	} {
		{ "1 second", 1000, nil },
		{ "2 seconds", 2000, nil },
		{ "1 hour 2 minutes and 1 second", 3721000, nil },	
		{ "1 hour 2 minutes 1 second", 3721000, nil },	
		{ "", 0, errors.New("Could not parse empty string into duration") },
		{ "duration", 0, errors.New("unexpected token 'duration'") },
		{ "1 hour 2", 0, errors.New("unexpected token '2'") },
		{ "hour hour", 0, errors.New("expected quantity, actually got 'hour'") },
		{ "1 lightyear", 0, errors.New("'lightyear', is not a valid unit") },
		{ "1 hour and 1 lightyear", 0, errors.New("'lightyear', is not a valid unit") },
		{ "1 hour 1 lightyear", 0, errors.New("'lightyear', is not a valid unit") },
	}

	for _, c := range cases {
		r, err := ParseDuration(c.Input)

		if c.Error != nil {
			if err == nil {
				t.Errorf("Expected error '%s' for input '%s'", err, c.Input)
			} else if err.Error() != c.Error.Error() {
				t.Errorf("Expected error '%s' for input '%s', actually got '%s'", c.Error, c.Input, err)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error '%s' for input '%s'", err, c.Input)
			}
			if c.Result != r {
				t.Errorf("Expected %d actual %d", c.Result, r)
			}
		}
	}
}
