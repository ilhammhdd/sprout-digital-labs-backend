package entity

import (
	"testing"

	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/errors"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/message"
	"github.com/stretchr/testify/assert"
)

func TestGetValidSquaresToMove_RookLightRightInit(t *testing.T) {
	givenRook := Rook{Indices: Indices{8, 8}}
	expectedValidSquares := Set[Square]{
		// vertical
		{'H', '2'}: {},
		{'H', '3'}: {},
		{'H', '4'}: {},
		{'H', '5'}: {},
		{'H', '6'}: {},
		{'H', '7'}: {},
		{'H', '8'}: {},
		// horizontal
		{'G', '1'}: {},
		{'F', '1'}: {},
		{'E', '1'}: {},
		{'D', '1'}: {},
		{'C', '1'}: {},
		{'B', '1'}: {},
		{'A', '1'}: {},
	}
	actualValidSquares := givenRook.GetValidSquaresToMove()
	assert.Equal(t, len(expectedValidSquares), len(actualValidSquares))
	for square := range expectedValidSquares {
		_, ok := actualValidSquares[square]
		assert.Truef(t, ok, "square: %s", square)
	}
}

func TestGetValidSquaresToMove_RookCenter(t *testing.T) {
	givenRook := Rook{Indices: Indices{5, 5}}
	expectedValidSquares := Set[Square]{
		// vertical
		{'E', '1'}: {},
		{'E', '2'}: {},
		{'E', '3'}: {},
		{'E', '5'}: {},
		{'E', '6'}: {},
		{'E', '7'}: {},
		{'E', '8'}: {},
		// horizontal
		{'H', '4'}: {},
		{'G', '4'}: {},
		{'F', '4'}: {},
		{'D', '4'}: {},
		{'C', '4'}: {},
		{'B', '4'}: {},
		{'A', '4'}: {},
	}
	actualValidSquares := givenRook.GetValidSquaresToMove()
	assert.Equal(t, len(expectedValidSquares), len(actualValidSquares))
	for square := range expectedValidSquares {
		_, ok := actualValidSquares[square]
		assert.Truef(t, ok, "square: %s", square)
	}
}

func TestGetDestDirection_RookInvalid(t *testing.T) {
	givenRook := Rook{Indices: Indices{3, 7}, Square: Square{'G', '6'}}
	actualDirs, actualErr := givenRook.GetDestDirection(Square{'E', '5'})
	assert.Nil(t, actualDirs)
	assert.Equal(t, errors.UnwrapTrace(errors.NewTrace(message.InvalidDirection)), errors.UnwrapTrace(actualErr))
}
