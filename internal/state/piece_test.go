package state

import (
	"testing"

	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestParsePiece_KingWhite(t *testing.T) {
	expectedPiece := entity.King{8, 5}
	actualPiece, actualErr := parsePiece(entity.Square{'E', '1'})
	assert.NoError(t, actualErr)
	assert.Equal(t, expectedPiece, actualPiece)
}

func TestParsePiece_QueenBlack(t *testing.T) {
	expectedPiece := entity.Queen{1, 4}
	actualPiece, actualErr := parsePiece(entity.Square{'D', '8'})
	assert.NoError(t, actualErr)
	assert.Equal(t, expectedPiece, actualPiece)
}

func TestParsePiece_RookWhiteRight(t *testing.T) {
	expectedPiece := entity.Rook{8, 8}
	actualPiece, actualErr := parsePiece(entity.Square{'H', '1'})
	assert.NoError(t, actualErr)
	assert.Equal(t, expectedPiece, actualPiece)
}

func TestParsePiece_KnightBlackLeft(t *testing.T) {
	expectedPiece := entity.Knight{1, 2}
	actualPiece, actualErr := parsePiece(entity.Square{'B', '8'})
	assert.NoError(t, actualErr)
	assert.Equal(t, expectedPiece, actualPiece)
}

func TestParsePiece_BishopWhiteRight(t *testing.T) {
	expectedPiece := entity.Bishop{Indices: entity.Indices{8, 6}, Square: entity.Square{'F', '1'}}
	actualPiece, actualErr := parsePiece(entity.Square{'F', '1'})
	assert.NoError(t, actualErr)
	assert.Equal(t, expectedPiece, actualPiece)
}

func TestParsePiece_PawnWhite(t *testing.T) {
	expectedPiece := entity.Pawn{2, 3}
	actualPiece, actualErr := parsePiece(entity.Square{'C', '7'})
	assert.NoError(t, actualErr)
	assert.Equal(t, expectedPiece, actualPiece)
}
