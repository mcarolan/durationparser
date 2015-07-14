package durationparser

import (
	"errors"
	"fmt"
	"strconv"
)

var unitInMillis = map[string]int64 {
	"milli": 1,
	"millis": 1,
	"second": 1000,
	"seconds": 1000,
	"minute": 60000,
	"minutes": 60000,
	"hour": 3600000,
	"hours": 3600000,
	"day": 86400000,
	"days": 86400000,
	"week": 604800000,
	"weeks": 604800000,
	"year": 31536000000,
	"years": 31536000000,
}

func ParseDuration(s string) (int64, error) {
	return parseDuration(tokenise(s))
}

func parseDuration(tokens []token) (int64, error) {
	if len(tokens) == 0 {
		return 0, nil
	} else if len(tokens) == 1 {
		return 0, errors.New(fmt.Sprintf("unexpected token '%s'", tokens[0].value))
	}

	quantityToken := tokens[0]
	unit := tokens[1]

	if quantityToken.tokenType != TOKEN_TYPE_NUMERIC {
		return 0, errors.New(fmt.Sprintf("expected quantity, actually got '%s'", quantityToken.value))
	}

	if unit.tokenType != TOKEN_TYPE_STRING {
		return 0, errors.New(fmt.Sprintf("expected unit, actually got '%s'", unit.value))
	}

	value, ok := unitInMillis[unit.value]

	if !ok {
		return 0, errors.New(fmt.Sprintf("'%s', is not a valid unit", unit.value))
	}

	quantity, _ := strconv.Atoi(quantityToken.value)

	result := int64(quantity) * value

	if len(tokens) > 2 {
		startIndex := 2
		if tokens[2].value == "and" {
			startIndex = 3
		}

		remaining, err := parseDuration(tokens[startIndex:])

		if err != nil {
			return 0, err
		}

		result += remaining
	}

	return result, nil
}
