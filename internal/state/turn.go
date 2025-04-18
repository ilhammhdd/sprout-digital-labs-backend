package state

type Turn struct {
	id          uint16
	playerWhite bool
	action      [2]string
}

var Turns []Turn
