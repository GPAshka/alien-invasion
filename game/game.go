package game

import (
	"alien-invasion/alien"
	"alien-invasion/world"
	"io"
)

type Game struct {
	alien.Aliens
	*world.Map
}

func NewGame(mapReader io.Reader, aliensNumber int) *Game {
	//create aliens
	aliens := alien.CreateAliens(aliensNumber)

	//create Map
	m := world.CreateMap(mapReader)

	game := &Game{Aliens: aliens, Map: m}
	return game
}
