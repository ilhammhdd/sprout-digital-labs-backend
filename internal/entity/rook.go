package entity

type Rook Indices

func (rook Rook) GetValidSquaresToMove() Set[Square] {
	validSquares := make(Set[Square])
	row, col := 1, rook[1]
	for row < 9 {
		if row != rook[0] {
			validSquares[getBoardSquare(Indices{row, col})] = NewSetElem()
		}
		row++
	}
	row, col = rook[0], 1
	for col < 9 {
		if col != rook[1] {
			validSquares[getBoardSquare(Indices{row, col})] = NewSetElem()
		}
		col++
	}
	return validSquares
}

func (rook Rook) GetDestDirection(dest Square) ([]Direction, error) {
	return nil, nil
}
