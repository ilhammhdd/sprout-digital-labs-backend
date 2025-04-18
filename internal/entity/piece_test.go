package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBoardSquare_TopLeft(t *testing.T) {
	givenIdx := Indices{1, 8}
	expectedSquare := Square{'H', '8'}
	actualSquare := getBoardSquare(givenIdx)
	assert.Equal(t, expectedSquare, actualSquare)
}

func TestGetBoardSquare_TopRight(t *testing.T) {
	givenIdx := Indices{1, 1}
	expectedSquare := Square{'A', '8'}
	actualSquare := getBoardSquare(givenIdx)
	assert.Equal(t, expectedSquare, actualSquare)
}

func TestGetBoardSquare_BottomRight(t *testing.T) {
	givenIdx := Indices{8, 8}
	expectedSquare := Square{'H', '1'}
	actualSquare := getBoardSquare(givenIdx)
	assert.Equal(t, expectedSquare, actualSquare)
}

func TestGetBoardSquare_BottomLeft(t *testing.T) {
	givenIdx := Indices{8, 1}
	expectedSquare := Square{'A', '1'}
	actualSquare := getBoardSquare(givenIdx)
	assert.Equal(t, expectedSquare, actualSquare)
}
