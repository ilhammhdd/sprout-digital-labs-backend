package usecase

import (
	"testing"

	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/entity"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/errors"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/message"
	"github.com/stretchr/testify/assert"
)

func TestValidateColumn_Uppercase(t *testing.T) {
	givenSquare := entity.Square{'H', '6'}
	actualErr := validateCol(givenSquare)
	assert.NoError(t, actualErr)
}

func TestValidateColumn_Lowercase(t *testing.T) {
	givenSquare := entity.Square{'e', '4'}
	actualErr := validateCol(givenSquare)
	assert.NoError(t, actualErr)
}

func TestValidateRow(t *testing.T) {
	givenSquare := entity.Square{'B', '4'}
	actualErr := validateRow(givenSquare)
	assert.NoError(t, actualErr)
}

func TestValidateCol_UppercaseOutOfBoundsUpper(t *testing.T) {
	givenSquare := entity.Square{'I', '5'}
	actualErr := validateCol(givenSquare)
	expectedErr := errors.NewTrace(message.SquareOutOfBounds)
	assert.Equal(t, expectedErr.Error(), actualErr.Error())
}

func TestValidateCol_UppercaseOutOfBoundsLower(t *testing.T) {
	givenSquare := entity.Square{64, '5'}
	actualErr := validateCol(givenSquare)
	expectedErr := errors.NewTrace(message.SquareOutOfBounds)
	assert.Equal(t, expectedErr.Error(), actualErr.Error())
}

func TestValidateCol_LowercaseOutOfBoundsUpper(t *testing.T) {
	givenSquare := entity.Square{'j', '5'}
	actualErr := validateCol(givenSquare)
	expectedErr := errors.NewTrace(message.SquareOutOfBounds)
	assert.Equal(t, expectedErr.Error(), actualErr.Error())
}

func TestValidateCol_LowercaseOutOfBoundsLower(t *testing.T) {
	givenSquare := entity.Square{63, '5'}
	actualErr := validateCol(givenSquare)
	expectedErr := errors.NewTrace(message.SquareOutOfBounds)
	assert.Equal(t, expectedErr.Error(), actualErr.Error())
}
func TestValidateRow_OutOfBoundsUpper(t *testing.T) {
	givenSquare := entity.Square{'C', '9'}
	actualErr := validateRow(givenSquare)
	expectedErr := errors.NewTrace(message.SquareOutOfBounds)
	assert.Equal(t, expectedErr.Error(), actualErr.Error())
}

func TestValidateRow_OutOfBoundsLower(t *testing.T) {
	givenSquare := entity.Square{'C', '0'}
	actualErr := validateRow(givenSquare)
	expectedErr := errors.NewTrace(message.SquareOutOfBounds)
	assert.Equal(t, expectedErr.Error(), actualErr.Error())
}
