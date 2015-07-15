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

func read(s string, offset int, pred readPred) string {
	var value bytes.Buffer

	for i := offset; i < len(s); i ++ {
		r := rune(s[i])
		if pred(r) {
			value.WriteRune(r)
		} else {
			break
		}
	}

	return value.String()
}


type readPred func(rune) bool

func tokenise(s string) []token {

	isOther := func(r rune) bool {
		return !(unicode.IsDigit(r) || unicode.IsLetter(r) || unicode.IsSpace(r))
	} 

	predicates := []struct {
		Predicate readPred
		TokenType tokenType
	} { 
		{ unicode.IsDigit, TOKEN_TYPE_NUMERIC },
		{ unicode.IsLetter, TOKEN_TYPE_STRING },
		{ isOther, TOKEN_TYPE_OTHER },
	}

	var result []token

	for i := 0; i < len(s); {
		//peek at this character
		r := rune(s[i])

		//skip if it's whitespace
		if unicode.IsSpace(r) {
			i += 1
			continue
		}

		//otherwise, consume as much of this type of character as possible and whack it in a token
		for _, p := range predicates {
			if (p.Predicate(r)) {
				value := read(s, i, p.Predicate)
				result = append(result, token { value, p.TokenType })
				//advance by the number of characters we read
				i += len(value)
				break
			}
		}
	}

	return result
}
