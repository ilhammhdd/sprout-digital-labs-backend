package entity

import (
	"testing"

	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/errors"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/message"
	"github.com/stretchr/testify/assert"
)

func TestGetValidSquaresToMove_BishopLeftDarkInit(t *testing.T) {
	givenBishop := Bishop{Indices: Indices{1, 3}}
	expectedValidSquares := Set[Square]{
		// diagonal bottom right
		{'D', '7'}: {},
		{'E', '6'}: {},
		{'F', '5'}: {},
		{'G', '4'}: {},
		{'H', '3'}: {},
		// diagonal bottom left
		{'B', '7'}: {},
		{'A', '6'}: {},
	}
	actualValidSquares := givenBishop.GetValidSquaresToMove()
	assert.Equal(t, len(expectedValidSquares), len(actualValidSquares))
	for square := range expectedValidSquares {
		_, ok := actualValidSquares[square]
		assert.Truef(t, ok, "square: %s", square)
	}
}

func TestGetValidSquaresToMove_BishopBottomLeft(t *testing.T) {
	givenBishop := Bishop{Indices: Indices{8, 1}}
	expectedValidSquares := Set[Square]{
		// diagonal top right
		{'B', '2'}: {},
		{'C', '3'}: {},
		{'D', '4'}: {},
		{'E', '5'}: {},
		{'F', '6'}: {},
		{'G', '7'}: {},
		{'H', '8'}: {},
	}
	actualValidSquares := givenBishop.GetValidSquaresToMove()
	assert.Equal(t, len(expectedValidSquares), len(actualValidSquares))
	for square := range expectedValidSquares {
		_, ok := actualValidSquares[square]
		assert.Truef(t, ok, "square: %s", square)
	}
}

func TestGetValidSquaresToMove_BishopCenter(t *testing.T) {
	givenBishop := Bishop{Indices: Indices{3, 6}}
	expectedValidSquares := Set[Square]{
		// diagonal top left
		{'E', '7'}: {},
		{'D', '8'}: {},
		// diagonal top right
		{'G', '7'}: {},
		{'H', '8'}: {},
		// diagonal bottom right
		{'G', '5'}: {},
		{'H', '4'}: {},
		// diagonal bottom left
		{'E', '5'}: {},
		{'D', '4'}: {},
		{'C', '3'}: {},
		{'B', '2'}: {},
		{'A', '1'}: {},
	}
	actualValidSquares := givenBishop.GetValidSquaresToMove()
	assert.Equal(t, len(expectedValidSquares), len(actualValidSquares))
	for square := range expectedValidSquares {
		_, ok := actualValidSquares[square]
		assert.Truef(t, ok, "square: %s", square)
	}
}

func TestGetDestDirection_BishopInvalid(t *testing.T) {
	givenBishop := Bishop{Indices: Indices{1, 8}, Square: Square{'H', '8'}}
	actualDirs, actualErr := givenBishop.GetDestDirection(Square{'H', '7'})
	assert.Equal(t, errors.UnwrapTrace(actualErr), errors.UnwrapTrace(errors.NewTrace(message.InvalidDirection)))
	assert.Nil(t, actualDirs)
}
