package usecase

import (
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/entity"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/errors"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/message"
)

func Move(origin, dest entity.Square) error {
	if err := validateSquare(origin, dest); err != nil {
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
