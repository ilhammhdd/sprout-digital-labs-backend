package usecase

import (
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/entity"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/errors"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/message"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/state"
)

func Move(origin, dest entity.Square) error {
	if err := validateSquare(origin, dest); err != nil {
		return err
	}
	origin, dest = colToUpperCase(origin), colToUpperCase(dest)
	piece, err := state.ParsePiece(origin)
	if err != nil {
		return err
	}
	validSquares := piece.GetValidSquaresToMove()
	if _, ok := validSquares[dest]; !ok {
		return errors.NewTrace(message.InvalidSquareToMove)
	}
	dirs, err := piece.GetDestDirection(dest)
	if err != nil {
		return err
	}
	if err := validateBetweenOriginAndDest(dirs, origin, dest); err != nil {
		return err
	}
	return nil
}

func validateSquare(origin, dest entity.Square) error {
	if err := validateCol(origin); err != nil {
		return err
	} else if err := validateCol(dest); err != nil {
		return err
	} else if err := validateRow(origin); err != nil {
		return err
	} else if err := validateRow(dest); err != nil {
		return err
	}
	return nil
}

func colToUpperCase(square entity.Square) entity.Square {
	if square[0] >= 'a' && square[0] <= 'h' {
		return entity.Square{square[0] - ('a' - 'A'), square[1]}
	}
	return square
}

func validateCol(coor entity.Square) error {
	if (coor[0] >= 'a' && coor[0] <= 'h') || (coor[0] >= 'A' && coor[0] <= 'H') {
		return nil
	}
	return errors.NewTrace(message.SquareOutOfBounds)
}

func validateRow(coor entity.Square) error {
	if coor[1] < '1' || coor[1] > '8' {
		return errors.NewTrace(message.SquareOutOfBounds)
	}
	return nil
}
