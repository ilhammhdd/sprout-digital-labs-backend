package entity

type Indices [2]int
type Square [2]byte

type Piece interface {
	GetValidSquaresToMove() Set[Square]
	GetDestDirection(dest Square) ([]Direction, error)
}

func TraverseDiagonal(idx Indices, rowOp, colOp int, fn func(row, col int)) {
	row, col := idx[0]+rowOp, idx[1]+colOp
	rowLim, colLim := 0, 0
	if rowOp == 1 {
		rowLim = 9
	}
	if colOp == 1 {
		colLim = 9
	}
	for row != rowLim && col != colLim {
		fn(row, col)
		row, col = row+rowOp, col+colOp
	}
}

func getBoardSquare(idx Indices) Square {
	return Square{byte(idx[1] + 'A' - 1), byte((9 - abs(idx[0])) + '0')}
}

func abs[T int | byte](a T) T {
	if a < 0 {
		return -a
	}
	return a
}
