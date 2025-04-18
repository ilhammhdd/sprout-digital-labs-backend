package entity

type Pawn Indices

func (pawn Pawn) GetValidSquaresToMove() Set[Square] {
	validSquares := make(Set[Square])
	if abs(pawn[0]) == 2 || pawn[0] == 7 {
		validSquares[pawn.getPawnBoardSquare(2)] = NewSetElem()
	}
	validSquares[pawn.getPawnBoardSquare(1)] = NewSetElem()
	return validSquares
}

func (pawn Pawn) getPawnBoardSquare(moves int) Square {
	if pawn[0] < 0 {
		return getBoardSquare(Indices{abs(pawn[0]) + moves, pawn[1]})
	}
	return getBoardSquare(Indices{pawn[0] - moves, pawn[1]})
}

func (pawn Pawn) GetDestDirection(dest Square) ([]Direction, error) {
	return nil, nil
}
