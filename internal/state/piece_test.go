package state

import (
	"testing"

	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestParsePiece_KingLight(t *testing.T) {
	expectedPiece := entity.King{Indices: entity.Indices{8, 5}, Square: entity.Square{'E', '1'}}
	actualPiece, actualErr := ParsePiece(entity.Square{'E', '1'})
	assert.NoError(t, actualErr)
	assert.Equal(t, expectedPiece, actualPiece)
}

func TestParsePiece_QueenDark(t *testing.T) {
	expectedPiece := entity.Queen{Indices: entity.Indices{1, 4}, Square: entity.Square{'D', '8'}}
	actualPiece, actualErr := ParsePiece(entity.Square{'D', '8'})
	assert.NoError(t, actualErr)
	assert.Equal(t, expectedPiece, actualPiece)
}

func TestParsePiece_RookLightRight(t *testing.T) {
	expectedPiece := entity.Rook{Indices: entity.Indices{8, 8}, Square: entity.Square{'H', '1'}}
	actualPiece, actualErr := ParsePiece(entity.Square{'H', '1'})
	assert.NoError(t, actualErr)
	assert.Equal(t, expectedPiece, actualPiece)
}

func TestParsePiece_KnightDarkLeft(t *testing.T) {
	expectedPiece := entity.Knight{Indices: entity.Indices{1, 2}, Square: entity.Square{'B', '8'}}
	actualPiece, actualErr := ParsePiece(entity.Square{'B', '8'})
	assert.NoError(t, actualErr)
	assert.Equal(t, expectedPiece, actualPiece)
}

func TestParsePiece_BishopLightRight(t *testing.T) {
	expectedPiece := entity.Bishop{Indices: entity.Indices{8, 6}, Square: entity.Square{'F', '1'}}
	actualPiece, actualErr := ParsePiece(entity.Square{'F', '1'})
	assert.NoError(t, actualErr)
	assert.Equal(t, expectedPiece, actualPiece)
}

func TestParsePiece_PawnLight(t *testing.T) {
	expectedPiece := entity.Pawn{Indices: entity.Indices{2, 3}, Square: entity.Square{'C', '7'}}
	actualPiece, actualErr := ParsePiece(entity.Square{'C', '7'})
	assert.NoError(t, actualErr)
	assert.Equal(t, expectedPiece, actualPiece)
}
