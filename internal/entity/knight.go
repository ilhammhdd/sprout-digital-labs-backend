package entity

type Knight Indices

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
	if knight[0]+rowOp > 0 && knight[0]+rowOp < 9 && knight[1]+colOp > 0 && knight[1]+colOp < 9 {
		fn(knight[0]+rowOp, knight[1]+colOp)
	}
}

func (knight Knight) GetDestDirection(dest Square) ([]Direction, error) {
	return nil, nil
}
