package entity

import (
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/errors"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/message"
)

type Rook struct {
	Indices
	Square
}

func (rook Rook) GetValidSquaresToMove() Set[Square] {
	validSquares := make(Set[Square])
	row, col := 1, rook.Indices[1]
	for row < 9 {
		if row != rook.Indices[0] {
			validSquares[getBoardSquare(Indices{row, col})] = NewSetElem()
		}
		row++
	}
	row, col = rook.Indices[0], 1
	for col < 9 {
		if col != rook.Indices[1] {
			validSquares[getBoardSquare(Indices{row, col})] = NewSetElem()
		}
		col++
	}
	return validSquares
}

func (rook Rook) GetDestDirection(dest Square) ([]Direction, error) {
	dir := getVerticalDirection(rook.Square, dest)
	if dir != None {
		return []Direction{dir}, nil
	}
	dir = getHorizontalDirection(rook.Square, dest)
	if dir != None {
		return []Direction{dir}, nil
	}
	return nil, errors.NewTrace(message.InvalidDirection)
}
