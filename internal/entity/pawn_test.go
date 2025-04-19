package entity

import (
	"testing"

	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/errors"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/message"
	"github.com/stretchr/testify/assert"
)

func TestGetValidSquaresToMove_PawnLightFirstMove(t *testing.T) {
	givenPawn := Pawn{Indices: Indices{7, 4}, Square: Square{'D', '2'}, IsLight: true}
	expectedValidSquares := Set[Square]{
		{'D', '3'}: {}, {'D', '4'}: {}, {'C', '3'}: {}, {'E', '3'}: {},
	}
	actualValidSquares := givenPawn.GetValidSquaresToMove()
	assert.Equal(t, len(expectedValidSquares), len(actualValidSquares))
	for square := range expectedValidSquares {
		_, ok := actualValidSquares[square]
		assert.Truef(t, ok, "square: %s", square)
	}
}

func TestGetValidSquaresToMove_PawnDarkFirstMove(t *testing.T) {
	givenPawn := Pawn{Indices: Indices{2, 5}, Square: Square{'E', '7'}, IsLight: false}
	expectedValidSquares := Set[Square]{
		{'E', '6'}: {}, {'E', '5'}: {}, {'D', '6'}: {}, {'F', '6'}: {},
	}
	actualValidSquares := givenPawn.GetValidSquaresToMove()
	assert.Equal(t, len(expectedValidSquares), len(actualValidSquares))
	for square := range expectedValidSquares {
		_, ok := actualValidSquares[square]
		assert.Truef(t, ok, "square: %s", square)
	}
}

func TestGetValidSquaresToMove_PawnLightRight(t *testing.T) {
	givenPawn := Pawn{Indices: Indices{3, 8}, Square: Square{'H', '6'}, IsLight: true}
	expectedValidSquares := Set[Square]{{'H', '7'}: {}, {'G', '7'}: {}}
	actualValidSquares := givenPawn.GetValidSquaresToMove()
	assert.Equal(t, len(expectedValidSquares), len(actualValidSquares))
	for square := range expectedValidSquares {
		_, ok := actualValidSquares[square]
		assert.Truef(t, ok, "square: %s", square)
	}
}

func TestGetValidSquaresToMove_PawnDarkCenter(t *testing.T) {
	givenPawn := Pawn{Indices: Indices{5, 5}, Square: Square{'E', '4'}, IsLight: false}
	expectedValidSquares := Set[Square]{{'E', '3'}: {}, {'D', '3'}: {}, {'F', '3'}: {}}
	actualValidSquares := givenPawn.GetValidSquaresToMove()
	assert.Equal(t, len(expectedValidSquares), len(actualValidSquares))
	for square := range expectedValidSquares {
		_, ok := actualValidSquares[square]
		assert.Truef(t, ok, "square: %s", square)
	}
}

func TestGetDestDirection_PawnCapture(t *testing.T) {
	givenPawn := Pawn{Indices: Indices{6, 1}, Square: Square{'A', '3'}}
	expectedDirs := []Direction{Top, Right}
	actualDirs, actualErr := givenPawn.GetDestDirection(Square{'B', '4'})
	assert.NoError(t, actualErr)
	assert.Equal(t, expectedDirs, actualDirs)
}

func TestGetDestDirection_PawnInvalid(t *testing.T) {
	givenPawn := Pawn{Indices: Indices{6, 1}, Square: Square{'A', '3'}}
	actualDirs, actualErr := givenPawn.GetDestDirection(Square{'B', '3'})
	assert.Nil(t, actualDirs)
	assert.Equal(t, errors.UnwrapTrace(errors.NewTrace(message.InvalidDirection)), errors.UnwrapTrace(actualErr))
}
