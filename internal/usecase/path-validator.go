package usecase

import (
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/entity"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/errors"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/message"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/state"
)

func validateBetweenOriginAndDest(dirs []entity.Direction, origin, dest entity.Square) error {
	oriIndices, destIndices := state.GetBoardIndices(origin), state.GetBoardIndices(dest)
	oriPiece, destPiece := state.Board[oriIndices[0]][oriIndices[1]], state.Board[destIndices[0]][destIndices[1]]
	if len(oriPiece) != 2 || oriPiece[1] == ' ' {
		return errors.NewTrace(message.NotAPiece)
	}
	validator := pathValidator{
		dirs:        dirs,
		oriIndices:  oriIndices,
		destIndices: destIndices,
		oriPiece:    oriPiece,
		destPiece:   destPiece,
	}
	if oriPiece[0] == 'T' {
		if destPiece == "  " || isNotSameColor(oriPiece, destPiece) {
			return nil
		}
		return errors.NewTrace(message.TheresAPieceBlocking)
	}
	return validator.validate()
}

type pathValidator struct {
	dirs        []entity.Direction
	oriIndices  entity.Indices
	destIndices entity.Indices
	oriPiece    string
	destPiece   string
}

func (validator pathValidator) validate() error {
	if len(validator.dirs) == 2 {
		return validator.validateDiagonal()
	}
	return validator.validateVerticalHorizontal()
}

func (validator pathValidator) validateDiagonal() error {
	if validator.dirs[0] == entity.Top && validator.dirs[1] == entity.Left {
		return validator.traverseDiagonal(-1, -1)
	} else if validator.dirs[0] == entity.Top && validator.dirs[1] == entity.Right {
		return validator.traverseDiagonal(-1, 1)
	} else if validator.dirs[0] == entity.Bottom && validator.dirs[1] == entity.Right {
		return validator.traverseDiagonal(1, 1)
	} else if validator.dirs[0] == entity.Bottom && validator.dirs[1] == entity.Left {
		return validator.traverseDiagonal(1, -1)
	}
	return nil
}

func (validator pathValidator) traverseDiagonal(rowOp, colOp int) error {
	row, col := validator.oriIndices[0]+rowOp, validator.oriIndices[1]+colOp
	var piece string
	for row != validator.destIndices[0] && col != validator.destIndices[1] {
		piece = state.Board[row][col]
		if piece != "  " {
			return errors.NewTrace(message.TheresAPieceBlocking)
		}
		row, col = row+rowOp, col+colOp
	}
	if isNotSameColor(validator.oriPiece, validator.destPiece) {
		return nil
	}
	return errors.NewTrace(message.TheresAPieceBlocking)
}

func (validator pathValidator) validateVerticalHorizontal() error {
	opOn, op := 0, 1
	if validator.dirs[0] == entity.Top || validator.dirs[0] == entity.Left {
		op = -1
	}
	if validator.dirs[0] == entity.Left || validator.dirs[0] == entity.Right {
		opOn = 1
	}
	rowCol := validator.oriIndices
	rowCol[opOn] += op
	for rowCol[opOn] != validator.destIndices[opOn] {
		if state.Board[rowCol[0]][rowCol[1]] != "  " {
			return errors.NewTrace(message.TheresAPieceBlocking)
		}
		rowCol[opOn] += op
	}
	if isNotSameColor(validator.oriPiece, validator.destPiece) {
		return nil
	}
	return errors.NewTrace(message.TheresAPieceBlocking)
}

func isNotSameColor(ori, dest string) bool {
	if len(ori) != 2 || len(dest) != 2 {
		return false
	}
	return ori[1] == 'L' && dest[1] == 'D' || ori[1] == 'D' && dest[1] == 'L' || dest == "  "
}
