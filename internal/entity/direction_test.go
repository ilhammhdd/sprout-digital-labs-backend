package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDiagonalDirection_TopLeft(t *testing.T) {
	expectedDirs := []Direction{Top, Left}
	actualDirs := getDiagonalDirection(Square{'H', '1'}, Square{'A', '8'})
	assert.Equal(t, expectedDirs, actualDirs)
}

func TestGetDiagonalDirection_TopRight(t *testing.T) {
	expectedDirs := []Direction{Top, Right}
	actualDirs := getDiagonalDirection(Square{'A', '1'}, Square{'H', '8'})
	assert.Equal(t, expectedDirs, actualDirs)
}

func TestGetDiagonalDirection_BottomRight(t *testing.T) {
	expectedDirs := []Direction{Bottom, Right}
	actualDirs := getDiagonalDirection(Square{'A', '8'}, Square{'H', '1'})
	assert.Equal(t, expectedDirs, actualDirs)
}

func TestGetDiagonalDirection_BottomLeft(t *testing.T) {
	expectedDirs := []Direction{Bottom, Left}
	actualDirs := getDiagonalDirection(Square{'H', '8'}, Square{'A', '1'})
	assert.Equal(t, expectedDirs, actualDirs)
}
func TestGetDiagonalDirection_Invalid(t *testing.T) {
	actualDirs := getDiagonalDirection(Square{'E', '4'}, Square{'C', '3'})
	assert.Nil(t, actualDirs)
}

func TestGetVerticalDirection_Top(t *testing.T) {
	assert.Equal(t, Top, getVerticalDirection(Square{'A', '1'}, Square{'A', '8'}))
}

func TestGetVerticalDirection_Bottom(t *testing.T) {
	assert.Equal(t, Bottom, getVerticalDirection(Square{'A', '8'}, Square{'A', '1'}))
}

func TestGetVerticalDirection_TopInvalid(t *testing.T) {
	assert.Equal(t, None, getVerticalDirection(Square{'A', '1'}, Square{'B', '8'}))
}

func TestGetVerticalDirection_BottomInvalid(t *testing.T) {
	assert.Equal(t, None, getVerticalDirection(Square{'A', '8'}, Square{'B', '1'}))
}

func TestGetHorizontalDirection_Left(t *testing.T) {
	assert.Equal(t, Left, getHorizontalDirection(Square{'G', '4'}, Square{'B', '4'}))
}

func TestGetHorizontalDirection_Right(t *testing.T) {
	assert.Equal(t, Right, getHorizontalDirection(Square{'A', '2'}, Square{'H', '2'}))
}

func TestGetHorizontalDirection_LeftInvalid(t *testing.T) {
	assert.Equal(t, None, getHorizontalDirection(Square{'G', '3'}, Square{'B', '4'}))
}

func TestGetHorizontalDirection_RightInvalid(t *testing.T) {
	assert.Equal(t, None, getHorizontalDirection(Square{'A', '2'}, Square{'H', '1'}))
}

// TODO: latest progress here
func TestGetLMoveDirection_LeftTop(t *testing.T) {
	assert.Equal(t, []Direction{Left, Top}, getLMoveDirection(Square{'C', '3'}, Square{'A', '4'}))
}

func TestGetLMoveDirection_TopLeft(t *testing.T) {
	assert.Equal(t, []Direction{Top, Left}, getLMoveDirection(Square{'C', '3'}, Square{'B', '5'}))
}

func TestGetLMoveDirection_TopRight(t *testing.T) {
	assert.Equal(t, []Direction{Top, Right}, getLMoveDirection(Square{'C', '3'}, Square{'D', '5'}))
}

func TestGetLMoveDirection_RightTop(t *testing.T) {
	assert.Equal(t, []Direction{Right, Top}, getLMoveDirection(Square{'C', '3'}, Square{'E', '4'}))
}

func TestGetLMoveDirection_RightBottom(t *testing.T) {
	assert.Equal(t, []Direction{Right, Bottom}, getLMoveDirection(Square{'C', '3'}, Square{'E', '2'}))
}

func TestGetLMoveDirection_BottomRight(t *testing.T) {
	assert.Equal(t, []Direction{Bottom, Right}, getLMoveDirection(Square{'C', '3'}, Square{'D', '1'}))
}

func TestGetLMoveDirection_BottomLeft(t *testing.T) {
	assert.Equal(t, []Direction{Bottom, Left}, getLMoveDirection(Square{'C', '3'}, Square{'B', '1'}))
}

func TestGetLMoveDirection_LeftBottom(t *testing.T) {
	assert.Equal(t, []Direction{Left, Bottom}, getLMoveDirection(Square{'C', '3'}, Square{'A', '2'}))
}
