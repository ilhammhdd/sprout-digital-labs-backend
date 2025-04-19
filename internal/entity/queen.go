package entity

import (
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/errors"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/message"
)

type Queen struct {
	Indices
	Square
}

func (queen Queen) GetValidSquaresToMove() Set[Square] {
	validSquares := make(Set[Square])
	setValidSquare := func(row, col int) {
		validSquares[getBoardSquare(Indices{row, col})] = NewSetElem()
	}

	row, col := queen.Indices[0], 1
	for col < 9 {
		if row != queen.Indices[0] || col != queen.Indices[1] {
			setValidSquare(row, col)
		}
		col++
	}

	row, col = 1, queen.Indices[1]
	for row < 9 {
		if row != queen.Indices[0] || col != queen.Indices[1] {
			setValidSquare(row, col)
		}
		row++
	}
	// diagonal top left
	TraverseDiagonal(queen.Indices, -1, -1, setValidSquare)
	// diagonal top right
	TraverseDiagonal(queen.Indices, -1, 1, setValidSquare)
	// diagonal bottom right
	TraverseDiagonal(queen.Indices, 1, 1, setValidSquare)
	// diagonal bottom left
	TraverseDiagonal(queen.Indices, 1, -1, setValidSquare)
	return validSquares
}

func (queen Queen) GetDestDirection(dest Square) ([]Direction, error) {
	dirs := getDiagonalDirection(queen.Square, dest)
	if dirs != nil {
		return dirs, nil
	}
	dir := getVerticalDirection(queen.Square, dest)
	if dir != None {
		return []Direction{dir}, nil
	}
	dir = getHorizontalDirection(queen.Square, dest)
	if dir != None {
		return []Direction{dir}, nil
	}
	return nil, errors.NewTrace(message.InvalidDirection)
}
