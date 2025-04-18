package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValidSquaresToMove_PawnWhiteFirstMove(t *testing.T) {
	givenPawn := Pawn{7, 4}
	expectedValidSquares := Set[Square]{{'D', '3'}: {}, {'D', '4'}: {}}
	actualValidSquares := givenPawn.GetValidSquaresToMove()
	assert.Equal(t, len(expectedValidSquares), len(actualValidSquares))
	for square := range expectedValidSquares {
		_, ok := actualValidSquares[square]
		assert.Truef(t, ok, "square: %s", square)
	}
}

func TestGetValidSquaresToMove_PawnBlackFirstMove(t *testing.T) {
	givenPawn := Pawn{-2, 5}
	expectedValidSquares := Set[Square]{{'E', '6'}: {}, {'E', '5'}: {}}
	actualValidSquares := givenPawn.GetValidSquaresToMove()
	assert.Equal(t, len(expectedValidSquares), len(actualValidSquares))
	for square := range expectedValidSquares {
		_, ok := actualValidSquares[square]
		assert.Truef(t, ok, "square: %s", square)
	}
}

func TestGetValidSquaresToMove_PawnWhiteCenter(t *testing.T) {
	givenPawn := Pawn{3, 8}
	expectedValidSquares := Set[Square]{{'H', '7'}: {}}
	actualValidSquares := givenPawn.GetValidSquaresToMove()
	assert.Equal(t, len(expectedValidSquares), len(actualValidSquares))
	for square := range expectedValidSquares {
		_, ok := actualValidSquares[square]
		assert.Truef(t, ok, "square: %s", square)
	}
}

func TestGetValidSquaresToMove_PawnBlackCenter(t *testing.T) {
	givenPawn := Pawn{-5, 5}
	expectedValidSquares := Set[Square]{{'E', '3'}: {}}
	actualValidSquares := givenPawn.GetValidSquaresToMove()
	assert.Equal(t, len(expectedValidSquares), len(actualValidSquares))
	for square := range expectedValidSquares {
		_, ok := actualValidSquares[square]
		assert.Truef(t, ok, "square: %s", square)
	}
}
