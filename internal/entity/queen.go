package entity

type Queen Indices

func (queen Queen) GetValidSquaresToMove() Set[Square] {
	validSquares := make(Set[Square])
	setValidSquare := func(row, col int) {
		validSquares[getBoardSquare(Indices{row, col})] = NewSetElem()
	}

	row, col := queen[0], 1
	for col < 9 {
		if row != queen[0] || col != queen[1] {
			setValidSquare(row, col)
		}
		col++
	}

	row, col = 1, queen[1]
	for row < 9 {
		if row != queen[0] || col != queen[1] {
			setValidSquare(row, col)
		}
		row++
	}
	// diagonal top left
	traverseDiagonal(Indices(queen), -1, -1, setValidSquare)
	// diagonal top right
	traverseDiagonal(Indices(queen), -1, 1, setValidSquare)
	// diagonal bottom right
	traverseDiagonal(Indices(queen), 1, 1, setValidSquare)
	// diagonal bottom left
	traverseDiagonal(Indices(queen), 1, -1, setValidSquare)
	return validSquares
}

func (queen Queen) GetDestDirection(dest Square) ([]Direction, error) {
	return nil, nil
}
