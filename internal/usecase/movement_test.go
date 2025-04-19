package usecase

import (
	"testing"

	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/entity"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/errors"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/pkg/message"
	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/state"
	"github.com/stretchr/testify/assert"
)

func TestColToUpperCase(t *testing.T) {
	assert.Equal(t, entity.Square{'F', '5'}, colToUpperCase(entity.Square{'f', '5'}))
}

func TestColToUpperCase_A(t *testing.T) {
	assert.Equal(t, entity.Square{'A', '5'}, colToUpperCase(entity.Square{'a', '5'}))
}

func TestColToUpperCase_H(t *testing.T) {
	assert.Equal(t, entity.Square{'H', '5'}, colToUpperCase(entity.Square{'h', '5'}))
}

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

var backedUpBoard [10][10]string

func TestValidateBetweenOriginAndDest_QueenBlockedHorizontal(t *testing.T) {
	backUpBoard()
	defer func() { restoreBoard() }()
	state.Board = [10][10]string{
		{" ", "A ", "B ", "C ", "D ", "E ", "F ", "G ", "H ", " "},
		{"8", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "8"},
		{"7", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "7"},
		{"6", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "6"},
		{"5", "  ", "  ", "  ", "  ", "QD", "BD", "  ", "  ", "5"},
		{"4", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "4"},
		{"3", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "3"},
		{"2", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "2"},
		{"1", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "1"},
		{" ", "A ", "B ", "C ", "D ", "E ", "F ", "G ", "H ", " "},
	}
	actualErr := validateBetweenOriginAndDest(
		[]entity.Direction{entity.Right}, entity.Square{'E', '5'}, entity.Square{'G', '5'},
	)
	assert.Equal(t, errors.UnwrapTrace(errors.NewTrace(message.TheresAPieceBlocking)), errors.UnwrapTrace(actualErr))
}

func TestValidateBetweenOriginAndDest_QueenBlockedVertical(t *testing.T) {
	backUpBoard()
	defer func() { restoreBoard() }()
	state.Board = [10][10]string{
		{" ", "A ", "B ", "C ", "D ", "E ", "F ", "G ", "H ", " "},
		{"8", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "8"},
		{"7", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "7"},
		{"6", "  ", "  ", "  ", "  ", "BD", "  ", "  ", "  ", "6"},
		{"5", "  ", "  ", "  ", "  ", "QD", "  ", "  ", "  ", "5"},
		{"4", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "4"},
		{"3", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "3"},
		{"2", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "2"},
		{"1", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "1"},
		{" ", "A ", "B ", "C ", "D ", "E ", "F ", "G ", "H ", " "},
	}
	actualErr := validateBetweenOriginAndDest(
		[]entity.Direction{entity.Top}, entity.Square{'E', '5'}, entity.Square{'E', '6'},
	)
	assert.Equal(t, errors.UnwrapTrace(errors.NewTrace(message.TheresAPieceBlocking)), errors.UnwrapTrace(actualErr))
}

func TestValidateBetweenOriginAndDest_QueenDiagonal(t *testing.T) {
	backUpBoard()
	defer func() { restoreBoard() }()
	state.Board = [10][10]string{
		{" ", "A ", "B ", "C ", "D ", "E ", "F ", "G ", "H ", " "},
		{"8", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "8"},
		{"7", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "7"},
		{"6", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "6"},
		{"5", "  ", "  ", "  ", "  ", "QD", "  ", "  ", "  ", "5"},
		{"4", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "4"},
		{"3", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "3"},
		{"2", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "2"},
		{"1", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "1"},
		{" ", "A ", "B ", "C ", "D ", "E ", "F ", "G ", "H ", " "},
	}
	actualErr := validateBetweenOriginAndDest(
		[]entity.Direction{entity.Top, entity.Right}, entity.Square{'E', '5'}, entity.Square{'H', '8'},
	)
	assert.NoError(t, actualErr)
}

func TestValidateBetweenOriginAndDest_QueenBlockedDiagonal(t *testing.T) {
	backUpBoard()
	defer func() { restoreBoard() }()
	state.Board = [10][10]string{
		{" ", "A ", "B ", "C ", "D ", "E ", "F ", "G ", "H ", " "},
		{"8", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "8"},
		{"7", "  ", "  ", "  ", "  ", "  ", "  ", "BD", "  ", "7"},
		{"6", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "6"},
		{"5", "  ", "  ", "  ", "  ", "QD", "  ", "  ", "  ", "5"},
		{"4", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "4"},
		{"3", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "3"},
		{"2", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "2"},
		{"1", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "1"},
		{" ", "A ", "B ", "C ", "D ", "E ", "F ", "G ", "H ", " "},
	}
	actualErr := validateBetweenOriginAndDest(
		[]entity.Direction{entity.Top, entity.Right}, entity.Square{'E', '5'}, entity.Square{'H', '8'},
	)
	assert.Equal(t, errors.UnwrapTrace(errors.NewTrace(message.TheresAPieceBlocking)), errors.UnwrapTrace(actualErr))
}

func TestValidateBetweenOriginAndDest_QueenCaptureDiagonal(t *testing.T) {
	backUpBoard()
	defer func() { restoreBoard() }()
	state.Board = [10][10]string{
		{" ", "A ", "B ", "C ", "D ", "E ", "F ", "G ", "H ", " "},
		{"8", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "BL", "8"},
		{"7", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "7"},
		{"6", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "6"},
		{"5", "  ", "  ", "  ", "  ", "QD", "  ", "  ", "  ", "5"},
		{"4", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "4"},
		{"3", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "3"},
		{"2", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "2"},
		{"1", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "1"},
		{" ", "A ", "B ", "C ", "D ", "E ", "F ", "G ", "H ", " "},
	}
	actualErr := validateBetweenOriginAndDest(
		[]entity.Direction{entity.Top, entity.Right}, entity.Square{'E', '5'}, entity.Square{'H', '8'},
	)
	assert.NoError(t, actualErr)
}

func TestValidateBetweenOriginAndDest_KnightBlocked(t *testing.T) {
	backUpBoard()
	defer func() { restoreBoard() }()
	state.Board = [10][10]string{
		{" ", "A ", "B ", "C ", "D ", "E ", "F ", "G ", "H ", " "},
		{"8", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "8"},
		{"7", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "7"},
		{"6", "  ", "  ", "  ", "  ", "  ", "  ", "BL", "  ", "6"},
		{"5", "  ", "  ", "  ", "  ", "TL", "  ", "  ", "  ", "5"},
		{"4", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "4"},
		{"3", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "3"},
		{"2", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "2"},
		{"1", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "1"},
		{" ", "A ", "B ", "C ", "D ", "E ", "F ", "G ", "H ", " "},
	}
	actualErr := validateBetweenOriginAndDest(
		[]entity.Direction{entity.Right, entity.Top}, entity.Square{'E', '5'}, entity.Square{'G', '6'},
	)
	assert.Equal(t, errors.UnwrapTrace(errors.NewTrace(message.TheresAPieceBlocking)), errors.UnwrapTrace(actualErr))
}

func TestValidateBetweenOriginAndDest_Knight(t *testing.T) {
	backUpBoard()
	defer func() { restoreBoard() }()
	state.Board = [10][10]string{
		{" ", "A ", "B ", "C ", "D ", "E ", "F ", "G ", "H ", " "},
		{"8", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "8"},
		{"7", "  ", "  ", "  ", "  ", "  ", "BD", "  ", "  ", "7"},
		{"6", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "6"},
		{"5", "  ", "  ", "  ", "  ", "TL", "  ", "  ", "  ", "5"},
		{"4", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "4"},
		{"3", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "3"},
		{"2", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "2"},
		{"1", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "1"},
		{" ", "A ", "B ", "C ", "D ", "E ", "F ", "G ", "H ", " "},
	}
	actualErr := validateBetweenOriginAndDest(
		[]entity.Direction{entity.Top, entity.Right}, entity.Square{'E', '5'}, entity.Square{'F', '7'},
	)
	assert.NoError(t, actualErr)
}

func backUpBoard() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			backedUpBoard[i][j] = state.Board[i][j]
		}
	}
}

func restoreBoard() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			state.Board[i][j] = backedUpBoard[i][j]
		}
	}
}
