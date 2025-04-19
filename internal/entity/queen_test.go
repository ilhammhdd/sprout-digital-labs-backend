package entity

import (
	"testing"

	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/errors"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/message"
	"github.com/stretchr/testify/assert"
)

func TestGetValidSquares_QueenDarkInit(t *testing.T) {
	givenQueen := Queen{Indices: Indices{1, 4}}
	expectedValidSquares := Set[Square]{
		// horizontal
		{'A', '8'}: {},
		{'B', '8'}: {},
		{'C', '8'}: {},
		{'E', '8'}: {},
		{'F', '8'}: {},
		{'G', '8'}: {},
		{'H', '8'}: {},
		// vertical
		{'D', '7'}: {},
		{'D', '6'}: {},
		{'D', '5'}: {},
		{'D', '4'}: {},
		{'D', '3'}: {},
		{'D', '2'}: {},
		{'D', '1'}: {},
		// diagonal bottom left
		{'C', '7'}: {},
		{'B', '6'}: {},
		{'A', '5'}: {},
		// diagonal bottom right
		{'E', '7'}: {},
		{'F', '6'}: {},
		{'G', '5'}: {},
		{'H', '4'}: {},
	}
	actualValidSquares := givenQueen.GetValidSquaresToMove()
	for square := range actualValidSquares {
		_, ok := expectedValidSquares[square]
		assert.Truef(t, ok, "square: %s", square)
	}
	assert.Equal(t, len(expectedValidSquares), len(actualValidSquares))
	for square := range expectedValidSquares {
		_, ok := actualValidSquares[square]
		assert.Truef(t, ok, "square: %s", square)
	}
}

func TestGetValidSquares_QueenTopLeft(t *testing.T) {
	givenQueen := Queen{Indices: Indices{1, 1}}
	expectedValidSquares := Set[Square]{
		// horizontal
		{'B', '8'}: {},
		{'C', '8'}: {},
		{'D', '8'}: {},
		{'E', '8'}: {},
		{'F', '8'}: {},
		{'G', '8'}: {},
		{'H', '8'}: {},
		// vertical
		{'A', '7'}: {},
		{'A', '6'}: {},
		{'A', '5'}: {},
		{'A', '4'}: {},
		{'A', '3'}: {},
		{'A', '2'}: {},
		{'A', '1'}: {},
		// diagonal bottom right
		{'B', '7'}: {},
		{'C', '6'}: {},
		{'D', '5'}: {},
		{'E', '4'}: {},
		{'F', '3'}: {},
		{'G', '2'}: {},
		{'H', '1'}: {},
	}
	actualValidSquares := givenQueen.GetValidSquaresToMove()
	for square := range actualValidSquares {
		_, ok := expectedValidSquares[square]
		assert.Truef(t, ok, "square: %s", square)
	}
	assert.Equal(t, len(expectedValidSquares), len(actualValidSquares))
	for square := range expectedValidSquares {
		_, ok := actualValidSquares[square]
		assert.Truef(t, ok, "square: %s", square)
	}
}

func TestGetValidSquares_QueenCenter(t *testing.T) {
	givenQueen := Queen{Indices: Indices{4, 4}}
	expectedValidSquares := Set[Square]{
		// horizontal
		{'A', '5'}: {},
		{'B', '5'}: {},
		{'C', '5'}: {},
		{'E', '5'}: {},
		{'F', '5'}: {},
		{'G', '5'}: {},
		{'H', '5'}: {},
		// vertical
		{'D', '8'}: {},
		{'D', '7'}: {},
		{'D', '6'}: {},
		{'D', '4'}: {},
		{'D', '3'}: {},
		{'D', '2'}: {},
		{'D', '1'}: {},
		// diagonal top left
		{'C', '6'}: {},
		{'B', '7'}: {},
		{'A', '8'}: {},
		// diagonal top right
		{'E', '6'}: {},
		{'F', '7'}: {},
		{'G', '8'}: {},
		// diagonal bottom right
		{'E', '4'}: {},
		{'F', '3'}: {},
		{'G', '2'}: {},
		{'H', '1'}: {},
		// diagonal bottom left
		{'C', '4'}: {},
		{'B', '3'}: {},
		{'A', '2'}: {},
	}
	actualValidSquares := givenQueen.GetValidSquaresToMove()
	for square := range actualValidSquares {
		_, ok := expectedValidSquares[square]
		assert.Truef(t, ok, "square: %s", square)
	}
	assert.Equal(t, len(expectedValidSquares), len(actualValidSquares))
	for square := range expectedValidSquares {
		_, ok := actualValidSquares[square]
		assert.Truef(t, ok, "square: %s", square)
	}
}

func TestGetDestDirection_QueenInvalid(t *testing.T) {
	givenQueen := Queen{Indices: Indices{3, 7}, Square: Square{'G', '6'}}
	actualDirs, actualErr := givenQueen.GetDestDirection(Square{'E', '5'})
	assert.Nil(t, actualDirs)
	assert.Equal(t, errors.UnwrapTrace(errors.NewTrace(message.InvalidDirection)), errors.UnwrapTrace(actualErr))
}
