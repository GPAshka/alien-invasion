package game

import (
	"alien-invasion/alien"
	"io"
)

type Game struct {
	alien.Aliens
	*Map
}

func NewGame(mapReader io.Reader, aliensNumber int) *Game {
	//create aliens
	aliens := alien.CreateAliens(aliensNumber)

	//create Map
	m := CreateMap(mapReader)

	game := &Game{Aliens: aliens, Map: m}
	return game
}
