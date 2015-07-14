package durationparser

import "testing"

func tokenArraysEqual(a []token, b []token) bool {
	if len(a) != len(b) {
		return false
	}

	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestTokenise(t *testing.T) {
	cases := [] struct {
		Input string
		Result []token
	} {
		{ Input: "second", Result: []token { token { "second", TOKEN_TYPE_STRING } } },	
		{ Input: "10 seconds", Result: []token { token { "10", TOKEN_TYPE_NUMERIC }, token { "seconds", TOKEN_TYPE_STRING }}},
		{ Input: "", Result: []token {}},
		{ Input: "Random(1 second, 10 seconds)", Result: []token {
			token { "Random", TOKEN_TYPE_STRING },
			token { "(", TOKEN_TYPE_OTHER },
			token { "1", TOKEN_TYPE_NUMERIC },
			token { "second", TOKEN_TYPE_STRING },
			token { ",", TOKEN_TYPE_OTHER },
			token { "10", TOKEN_TYPE_NUMERIC },
			token { "seconds", TOKEN_TYPE_STRING },
			token { ")", TOKEN_TYPE_OTHER },
		}},
	}

	for _, c := range cases {
		r := tokenise(c.Input)

		if !tokenArraysEqual(c.Result, r) {
			t.Errorf("Expected %v actual %v", c.Result, r)
		}
	}
}
