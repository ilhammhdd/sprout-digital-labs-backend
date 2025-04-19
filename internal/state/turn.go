package state

import "github.com/ilhammhdd/sprout-digital-labs-backend/internal/entity"

type Turn struct {
	ID       uint16
	IsLight  bool
	Movement [2]entity.Square
}

var Turns []Turn
