package entity

import (
	"testing"

	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/errors"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/message"
	"github.com/stretchr/testify/assert"
)

func TestGetValidSquaresToMove_KnightLeftLightInit(t *testing.T) {
	givenKnight := Knight{Indices: Indices{8, 2}}
	expectedValidSquares := Set[Square]{
		{'C', '3'}: {},
		{'D', '2'}: {},
		{'A', '3'}: {},
	}
	actualValidSquares := givenKnight.GetValidSquaresToMove()
	assert.Equal(t, len(expectedValidSquares), len(actualValidSquares))
	for square := range expectedValidSquares {
		_, ok := actualValidSquares[square]
		assert.Truef(t, ok, "square: %s", square)
	}
}

func TestGetValidSquaresToMove_KnightCenter(t *testing.T) {
	givenKnight := Knight{Indices: Indices{4, 4}}
	expectedValidSquares := Set[Square]{
		{'B', '6'}: {},
		{'C', '7'}: {},
		{'E', '7'}: {},
		{'F', '6'}: {},
		{'F', '4'}: {},
		{'E', '3'}: {},
		{'C', '3'}: {},
		{'B', '4'}: {},
	}
	actualValidSquares := givenKnight.GetValidSquaresToMove()
	assert.Equal(t, len(expectedValidSquares), len(actualValidSquares))
	for square := range expectedValidSquares {
		_, ok := actualValidSquares[square]
		assert.Truef(t, ok, "square: %s", square)
	}
}

func TestGetDestDirection_KnightInvalid(t *testing.T) {
	givenKnight := Knight{Indices: Indices{5, 5}, Square: Square{'E', '5'}}
	actualDir, actualErr := givenKnight.GetDestDirection(Square{'C', '7'})
	assert.Nil(t, actualDir)
	assert.Equal(t, errors.UnwrapTrace(errors.NewTrace(message.InvalidDirection)), errors.UnwrapTrace(actualErr))
}
