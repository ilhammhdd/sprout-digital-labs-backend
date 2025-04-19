package entity

const (
	None   Direction = -1
	Left   Direction = 0
	Top    Direction = 1
	Right  Direction = 2
	Bottom Direction = 3
)

type Direction int8

func getDiagonalDirection(origin, dest Square) []Direction {
	if max(origin[0], dest[0])-min(origin[0], dest[0]) != max(origin[1], dest[1])-min(origin[1], dest[1]) {
		return nil
	}
	if origin[0] > dest[0] && origin[1] < dest[1] {
		return []Direction{Top, Left}
	} else if origin[0] < dest[0] && origin[1] < dest[1] {
		return []Direction{Top, Right}
	} else if origin[0] < dest[0] && origin[1] > dest[1] {
		return []Direction{Bottom, Right}
	} else if origin[0] > dest[0] && origin[1] > dest[1] {
		return []Direction{Bottom, Left}
	}
	return nil
}

func getVerticalDirection(origin, dest Square) Direction {
	if origin[0] != dest[0] {
		return None
	}
	if origin[1] < dest[1] {
		return Top
	} else if origin[1] > dest[1] {
		return Bottom
	}
	return None
}

func getHorizontalDirection(origin, dest Square) Direction {
	if origin[1] != dest[1] {
		return None
	}
	if origin[0] > dest[0] {
		return Left
	} else if origin[0] < dest[0] {
		return Right
	}
	return None
}

func getLMoveDirection(origin, dest Square) []Direction {
	if origin[0] == dest[0]+2 && origin[1] == dest[1]-1 {
		return []Direction{Left, Top}
	} else if origin[0] == dest[0]+1 && origin[1] == dest[1]-2 {
		return []Direction{Top, Left}
	} else if origin[0] == dest[0]-1 && origin[1] == dest[1]-2 {
		return []Direction{Top, Right}
	} else if origin[0] == dest[0]-2 && origin[1] == dest[1]-1 {
		return []Direction{Right, Top}
	} else if origin[0] == dest[0]-2 && origin[1] == dest[1]+1 {
		return []Direction{Right, Bottom}
	} else if origin[0] == dest[0]-1 && origin[1] == dest[1]+2 {
		return []Direction{Bottom, Right}
	} else if origin[0] == dest[0]+1 && origin[1] == dest[1]+2 {
		return []Direction{Bottom, Left}
	} else if origin[0] == dest[0]+2 && origin[1] == dest[1]+1 {
		return []Direction{Left, Bottom}
	}
	return nil
}
