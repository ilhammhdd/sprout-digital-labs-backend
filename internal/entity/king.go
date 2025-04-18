package entity

type King Indices

func (king King) GetValidSquaresToMove() Set[Square] {
	validSquares := make(Set[Square])
	rowLim, colLim := king[0]+1, king[1]+1
	row := king[0] - 1
	for row <= rowLim {
		col := king[1] - 1
		for col <= colLim {
			if (row > 0 && row < 9 && col > 0 && col < 9) && (row != king[0] || col != king[1]) {
				validSquares[getBoardSquare(Indices{row, col})] = NewSetElem()
			}
			col++
		}
		row++
	}
	return validSquares
}

func (king King) GetDestDirection(dest Square) ([]Direction, error) {
	return nil, nil
}
