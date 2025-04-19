package entity

import (
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/errors"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/message"
)

type Pawn struct {
	Indices
	Square
	IsLight bool
}

func (pawn Pawn) GetValidSquaresToMove() Set[Square] {
	validSquares := make(Set[Square])
	if (pawn.IsLight && pawn.Indices[0] == 7) || pawn.Indices[0] == 2 {
		validSquares[pawn.getPawnBoardSquare(2)] = NewSetElem()
	}
	validSquares[pawn.getPawnBoardSquare(1)] = NewSetElem()
	if pawn.IsLight && pawn.Square[1]+1 < '9' {
		if pawn.Square[0]+1 < 'I' {
			validSquares[Square{pawn.Square[0] + 1, pawn.Square[1] + 1}] = NewSetElem()
		}
		if pawn.Square[0]-1 >= 'A' {
			validSquares[Square{pawn.Square[0] - 1, pawn.Square[1] + 1}] = NewSetElem()
		}
	} else if pawn.Square[1]-1 > '0' {
		if pawn.Square[0]-1 >= 'A' {
			validSquares[Square{pawn.Square[0] - 1, pawn.Square[1] - 1}] = NewSetElem()
		}
		if pawn.Square[0]+1 < 'I' {
			validSquares[Square{pawn.Square[0] + 1, pawn.Square[1] - 1}] = NewSetElem()
		}
	}
	return validSquares
}

func (pawn Pawn) getPawnBoardSquare(moves int) Square {
	if !pawn.IsLight {
		return getBoardSquare(Indices{pawn.Indices[0] + moves, pawn.Indices[1]})
	}
	return getBoardSquare(Indices{pawn.Indices[0] - moves, pawn.Indices[1]})
}

func (pawn Pawn) GetDestDirection(dest Square) ([]Direction, error) {
	dirs := getDiagonalDirection(pawn.Square, dest)
	if dirs != nil {
		return dirs, nil
	}
	dir := getVerticalDirection(pawn.Square, dest)
	if dir != None {
		return []Direction{dir}, nil
	}
	return nil, errors.NewTrace(message.InvalidDirection)
}
