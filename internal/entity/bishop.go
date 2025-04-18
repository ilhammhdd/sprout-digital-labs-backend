package entity

import (
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/errors"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/message"
)

type Bishop struct {
	Indices
	Square
}

func (bishop Bishop) GetValidSquaresToMove() Set[Square] {
	validSquares := make(Set[Square])
	setValidSquares := func(row, col int) {
		validSquares[getBoardSquare(Indices{row, col})] = NewSetElem()
	}
	// diagonal top left
	traverseDiagonal(bishop.Indices, -1, -1, setValidSquares)
	// diagonal top right
	traverseDiagonal(bishop.Indices, -1, 1, setValidSquares)
	// diagonal bottom right
	traverseDiagonal(bishop.Indices, 1, 1, setValidSquares)
	// diagonal bottom left
	traverseDiagonal(bishop.Indices, 1, -1, setValidSquares)
	return validSquares
}

func (bishop Bishop) GetDestDirection(dest Square) ([]Direction, error) {
	dirs := getDiagonalDirection(bishop.Square, dest)
	if len(dirs) == 0 {
		return nil, errors.NewTrace(message.InvalidDirection)
	}
	return dirs, nil
}
