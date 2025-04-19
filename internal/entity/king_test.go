package entity

import (
	"testing"

	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/errors"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/message"
	"github.com/stretchr/testify/assert"
)

func TestGetValidSquares_KingLightInit(t *testing.T) {
	givenKing := King{Indices: Indices{8, 5}}
	expectedValidSquares := Set[Square]{
		{'D', '1'}: {},
		{'D', '2'}: {},
		{'E', '2'}: {},
		{'F', '2'}: {},
		{'F', '1'}: {},
	}
	actualValidSquares := givenKing.GetValidSquaresToMove()
	assert.Equal(t, len(expectedValidSquares), len(actualValidSquares))
	for square := range expectedValidSquares {
		_, ok := actualValidSquares[square]
		assert.Truef(t, ok, "square: %s", square)
	}
}

func TestGetValidSquares_KingCenter(t *testing.T) {
	givenKing := King{Indices: Indices{4, 4}}
	expectedValidSquares := Set[Square]{
		{'C', '6'}: {},
		{'D', '6'}: {},
		{'E', '6'}: {},
		{'C', '5'}: {},
		{'E', '5'}: {},
		{'C', '4'}: {},
		{'D', '4'}: {},
		{'E', '4'}: {},
	}
	actualValidSquares := givenKing.GetValidSquaresToMove()
	assert.Equal(t, len(expectedValidSquares), len(actualValidSquares))
	for square := range expectedValidSquares {
		_, ok := actualValidSquares[square]
		assert.Truef(t, ok, "square: %s", square)
	}
}

func TestGetDestDirection_KingInvalid(t *testing.T) {
	givenKing := King{Indices: Indices{5, 5}, Square: Square{'E', '4'}}
	actualDir, err := givenKing.GetDestDirection(Square{'C', '3'})
	assert.Nil(t, actualDir)
	assert.Equal(t, errors.UnwrapTrace(errors.NewTrace(message.InvalidDirection)), errors.UnwrapTrace(err))
}
