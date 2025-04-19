package entity

import (
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/errors"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/message"
)

type King struct {
	Indices
	Square
}

func (king King) GetValidSquaresToMove() Set[Square] {
	validSquares := make(Set[Square])
	rowLim, colLim := king.Indices[0]+1, king.Indices[1]+1
	row := king.Indices[0] - 1
	for row <= rowLim {
		col := king.Indices[1] - 1
		for col <= colLim {
			if (row > 0 && row < 9 && col > 0 && col < 9) && (row != king.Indices[0] || col != king.Indices[1]) {
				validSquares[getBoardSquare(Indices{row, col})] = NewSetElem()
			}
			col++
		}
		row++
	}
	return validSquares
}

func (king King) GetDestDirection(dest Square) ([]Direction, error) {
	dir := getHorizontalDirection(king.Square, dest)
	if dir != None {
		return []Direction{dir}, nil
	}
	dir = getVerticalDirection(king.Square, dest)
	if dir != None {
		return []Direction{dir}, nil
	}
	dirs := getDiagonalDirection(king.Square, dest)
	if dirs != nil {
		return dirs, nil
	}
	return nil, errors.NewTrace(message.InvalidDirection)
}
