package durationparser

import (
	"unicode"
	"bytes"
)

type tokenType int

const (
	TOKEN_TYPE_STRING tokenType = iota
	TOKEN_TYPE_NUMERIC
	TOKEN_TYPE_OTHER
)

type token struct {
	value string
	tokenType tokenType
}

type readPred func(rune) bool

func read(s string, offset int, pred readPred) string {
	var value bytes.Buffer

	for x := 0; (offset + x) < len(s); x ++ {
		c := s[offset + x]
		if pred(rune(c)) {
			value.WriteString(string(c))
		} else {
			break
		}
	}

	return value.String()
}


func tokenise(s string) []token {
	var result []token

	predicates := []struct {
		Predicate readPred
		TokenType tokenType
	} { 
		{ unicode.IsDigit, TOKEN_TYPE_NUMERIC },
		{ unicode.IsLetter, TOKEN_TYPE_STRING },
	}

	for i := 0; i < len(s); {
		r := rune(s[i])
		if unicode.IsSpace(r) {
			i += 1
			continue
		}

		tokenRead := false

		for _, p := range predicates {
			if (p.Predicate(r)) {
				value := read(s, i, p.Predicate)
				result = append(result, token { value, p.TokenType })
				i += len(value)
				tokenRead = true
				break
			}
		}

		if !tokenRead {
			result = append(result, token { string(r), TOKEN_TYPE_OTHER })
			i += 1
		}
	}

	return result
}
