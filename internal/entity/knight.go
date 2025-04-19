package entity

import (
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/errors"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/message"
)

type Knight struct {
	Indices
	Square
}

func (knight Knight) GetValidSquaresToMove() Set[Square] {
	validSquares := make(Set[Square])
	setValidSquares := func(row, col int) {
		validSquares[getBoardSquare(Indices{row, col})] = NewSetElem()
	}

	// 2 top left
	knight.setIfInBoundary(-1, -2, setValidSquares)
	knight.setIfInBoundary(-2, -1, setValidSquares)
	// 2 top right
	knight.setIfInBoundary(-2, 1, setValidSquares)
	knight.setIfInBoundary(-1, 2, setValidSquares)
	// 2 bottom right
	knight.setIfInBoundary(1, 2, setValidSquares)
	knight.setIfInBoundary(2, 1, setValidSquares)
	// 2 bottom left
	knight.setIfInBoundary(2, -1, setValidSquares)
	knight.setIfInBoundary(1, -2, setValidSquares)

	return validSquares
}

func (knight Knight) setIfInBoundary(rowOp, colOp int, fn func(row, col int)) {
	if knight.Indices[0]+rowOp > 0 && knight.Indices[0]+rowOp < 9 && knight.Indices[1]+colOp > 0 && knight.Indices[1]+colOp < 9 {
		fn(knight.Indices[0]+rowOp, knight.Indices[1]+colOp)
	}
}

func (knight Knight) GetDestDirection(dest Square) ([]Direction, error) {
	dirs := getLMoveDirection(knight.Square, dest)
	if dirs == nil {
		return nil, errors.NewTrace(message.InvalidDirection)
	}
	return dirs, nil
}
