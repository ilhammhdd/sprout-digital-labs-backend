package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValidSquaresToMove_KnightLeftWhiteInit(t *testing.T) {
	givenKnight := Knight{8, 2}
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
	givenKnight := Knight{4, 4}
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
