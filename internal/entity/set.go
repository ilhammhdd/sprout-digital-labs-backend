package entity

type SetElem struct{}
type Set[T comparable] map[T]SetElem

func NewSetElem() SetElem {
	return SetElem{}
}
